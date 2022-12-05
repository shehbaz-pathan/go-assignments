package main

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddBookDB(t *testing.T) {
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
		AddRow(9780772013040, "The Suicide Murders", "Book by Engel, Howard", "ENGEL, Howard", 62.92).
		AddRow(9781982156909, "The Comedy of Errors", "The authoritative edition of The Comedy of Errors from The Folger Shakespeare Library, the trusted and widely used Shakespeare series for students and general readers", "William Shakespeare", 10.39)
	mock.ExpectPrepare("select (.+) from book_info")
	mock.ExpectQuery("select (.+) from book_info").WillReturnRows(rows)
	expectedResult := []Book{
		{Isbn: 9780772013040, Title: "The Suicide Murders", Synopsis: "Book by Engel, Howard", AuthorName: "ENGEL, Howard", Price: 62.92},
		{Isbn: 9781982156909, Title: "The Comedy of Errors", Synopsis: "The authoritative edition of The Comedy of Errors from The Folger Shakespeare Library, the trusted and widely used Shakespeare series for students and general readers", AuthorName: "William Shakespeare", Price: 10.39},
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
	titles := []string{"Test Book", "Fake Book", "Best Book"}
	for i, title := range titles {
		rows := sqlmock.NewRows([]string{"Isbn", "Title", "Synopsis", "AuthorName", "Price"}).
			AddRow(123+i, title, "This is a "+title, "Unkown Author", 120.00+float32(i))
		mock.ExpectPrepare("select (.+) from book_info where Title = ?")
		mock.ExpectQuery("select (.+) from book_info where Title = ?").WithArgs(title).WillReturnRows(rows)
		expectedResult := []Book{
			{Isbn: 123 + int64(i), Title: title, Synopsis: "This is a " + title, AuthorName: "Unkown Author", Price: 120.00 + float32(i)},
		}
		Result, err := dbops.GetBookByTitle(title)
		assert.Equal(t, nil, err)
		assert.Equal(t, expectedResult, Result)
	}
}
