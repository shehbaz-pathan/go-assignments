package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	rt := gin.Default()
	return rt
}

func TestAddBookHandler(t *testing.T) {
	book := Book{
		Isbn:       9780772013040,
		Title:      "The Suicide Murders",
		Synopsis:   "Book by Engel, Howard",
		AuthorName: "ENGEL, Howard",
		Price:      62.92,
	}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Error %s while opening mock DB connection", err)
	}
	defer db.Close()
	dbops := &DBHandler{db}
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("insert into book_info(Isbn,Title,Synopsis,AuthorName,Price) values(?,?,?,?,?)")
	mock.ExpectExec("insert into book_info(Isbn,Title,Synopsis,AuthorName,Price) values(?,?,?,?,?)").WithArgs(book.Isbn, book.Title, book.Synopsis, book.AuthorName, book.Price).WillReturnResult(result)
	r := SetupRouter()
	r.POST("/addbook", AddBookHandler(dbops))
	Data, err := json.Marshal(book)
	if err != nil {
		t.Fatalf("Error %s while json marshaling", err)
	}
	req, _ := http.NewRequest(http.MethodPost, "/addbook", bytes.NewBuffer(Data))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, book, book)
}
func TestGetBooksHandler(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error %s while opening mock DB connection", err)
	}
	defer db.Close()
	dbops := &DBHandler{db}
	rows := sqlmock.NewRows([]string{"Isbn", "Title", "Synopsis", "AuthorName", "Price"}).
		AddRow(9780772013040, "The Suicide Murders", "Book by Engel, Howard", "ENGEL, Howard", 62.92).
		AddRow(9781982156909, "The Comedy of Errors", "The authoritative edition of The Comedy of Errors from The Folger Shakespeare Library, the trusted and widely used Shakespeare series for students and general readers", "William Shakespeare", 10.39)
	mock.ExpectPrepare("select (.+) from book_info")
	mock.ExpectQuery("select (.+) from book_info").WillReturnRows(rows)
	r := SetupRouter()
	r.GET("/getbooks", GetBooksHandler(dbops))
	req, _ := http.NewRequest(http.MethodGet, "/getbooks", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	expectedResult := []Book{
		{Isbn: 9780772013040, Title: "The Suicide Murders", Synopsis: "Book by Engel, Howard", AuthorName: "ENGEL, Howard", Price: 62.92},
		{Isbn: 9781982156909, Title: "The Comedy of Errors", Synopsis: "The authoritative edition of The Comedy of Errors from The Folger Shakespeare Library, the trusted and widely used Shakespeare series for students and general readers", AuthorName: "William Shakespeare", Price: 10.39},
	}
	Result := []Book{}
	json.Unmarshal(w.Body.Bytes(), &Result)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expectedResult, Result)
}
func TestGetBookByTitleHandler(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error %s while opening mock DB connection", err)
	}
	defer db.Close()
	dbops := &DBHandler{db}
	rows := sqlmock.NewRows([]string{"Isbn", "Title", "Synopsis", "AuthorName", "Price"}).
		AddRow(9781982156909, "The Comedy of Errors", "The authoritative edition of The Comedy of Errors from The Folger Shakespeare Library, the trusted and widely used Shakespeare series for students and general readers", "William Shakespeare", 10.39)
	mock.ExpectPrepare("select (.+) from book_info where Title = ?")
	mock.ExpectQuery("select (.+) from book_info where Title = ?").WithArgs("The Comedy of Errors").WillReturnRows(rows)
	r := SetupRouter()
	r.GET("/getbookbytitle/:title", GetBookByTitleHandler(dbops))
	req, _ := http.NewRequest(http.MethodGet, "/getbookbytitle/The Comedy of Errors", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	expectedResult := []Book{
		{Isbn: 9781982156909, Title: "The Comedy of Errors", Synopsis: "The authoritative edition of The Comedy of Errors from The Folger Shakespeare Library, the trusted and widely used Shakespeare series for students and general readers", AuthorName: "William Shakespeare", Price: 10.39},
	}
	Result := []Book{}
	json.Unmarshal(w.Body.Bytes(), &Result)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expectedResult, Result)
	/*data := map[string]string {"Test Book", "Fake Book", "Latest Book"}
	for i, title := range titles {
		rows := sqlmock.NewRows([]string{"Isbn", "Title", "Synopsis", "AuthorName", "Price"}).
			AddRow(123+i, title, "This is a "+title, "Shehbaz", 120.00+float32(i))
		mock.ExpectPrepare("select (.+) from book_info where Title = ?")
		mock.ExpectQuery("select (.+) from book_info where Title = ?").WithArgs(title).WillReturnRows(rows)
		r := SetupRouter()
		r.GET("/getbookbytitle/:title", GetBookByTitleHandler(dbops))
		path := "/getbookbytitle/" + title
		req, _ := http.NewRequest(http.MethodGet, path, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		expectedResult := []Book{
			{Isbn: 123 + int64(i), Title: title, Synopsis: "This is a " + title, AuthorName: "Shehbaz", Price: 120.00 + float32(i)},
		}
		Result := []Book{}
		json.Unmarshal(w.Body.Bytes(), &Result)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, expectedResult, Result)
	}*/
}
