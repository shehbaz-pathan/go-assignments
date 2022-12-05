package main

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddBookDB(t *testing.T) {
	book := Book{
		Isbn:       123,
		Title:      "Fake Book",
		Synopsis:   "This is fake book",
		AuthorName: "Shehbaz",
		Price:      100.20,
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
	err = dbops.AddBook(book)
	assert.Equal(t, nil, err)
}

func TestGetBooksDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error %s while opening mock DB connection", err)
	}
	defer db.Close()
	dbops := &DBHandler{db}
	rows := sqlmock.NewRows([]string{"Isbn", "Title", "Synopsis", "AuthorName", "Price"}).
		AddRow(123, "Test Book", "This is a test book", "Shehbaz", 120.00).
		AddRow(124, "Fake Book", "This is a fake book", "Shahzain", 150.50)
	mock.ExpectPrepare("select (.+) from book_info")
	mock.ExpectQuery("select (.+) from book_info").WillReturnRows(rows)
	expectedResult := []Book{
		{Isbn: 123, Title: "Test Book", Synopsis: "This is a test book", AuthorName: "Shehbaz", Price: 120.00},
		{Isbn: 124, Title: "Fake Book", Synopsis: "This is a fake book", AuthorName: "Shahzain", Price: 150.50},
	}
	Result, err := dbops.GetBooks()
	assert.Equal(t, nil, err)
	assert.Equal(t, expectedResult, Result)
}

func TestGetBookByTitleDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error %s while opening mock DB connection", err)
	}
	defer db.Close()
	dbops := &DBHandler{db}
	titles := []string{"Test Book", "Fake Book", "Latest Book"}
	for i, title := range titles {
		rows := sqlmock.NewRows([]string{"Isbn", "Title", "Synopsis", "AuthorName", "Price"}).
			AddRow(123+i, title, "This is a "+title, "Shehbaz", 120.00+float32(i))
		mock.ExpectPrepare("select (.+) from book_info where Title = ?")
		mock.ExpectQuery("select (.+) from book_info where Title = ?").WithArgs(title).WillReturnRows(rows)
		expectedResult := []Book{
			{Isbn: 123 + int64(i), Title: title, Synopsis: "This is a " + title, AuthorName: "Shehbaz", Price: 120.00 + float32(i)},
		}
		Result, err := dbops.GetBookByTitle(title)
		assert.Equal(t, nil, err)
		assert.Equal(t, expectedResult, Result)
	}
}
