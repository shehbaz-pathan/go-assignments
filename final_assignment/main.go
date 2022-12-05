package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

type DBOperations interface {
	AddBook(b Book) error
	GetBooks() ([]Book,error)
	GetBookByTitle(s string) ([]Book,error)
}

type Book struct {
	Isbn int64 `json:"isbn" form:"isbn" binding:"required,gte=1"`
	Title string `json:"title" form:"title" binding:"required,min=1"`
	Synopsis string `json:"synopsis" form:"synopsis" binding:"required,min=5"`
	AuthorName string `json:"authorname" form:"authorname" binding:"required,min=2"`
	Price float32 `json:"price" form:"price" binding:"required"`
}

type DBHandler struct {
	db *sql.DB
}
func main() {
	db,err:=DBConnection()
	if err!=nil {
		fmt.Println("Error connecting DB",err)
		return
	}
	defer db.Close()
	dbops:= &DBHandler{db}
	router:= gin.Default()
	router.LoadHTMLFiles("./templates/index.html")
	router.GET("/getbooks",GetBooksHandler(dbops))
	router.POST("/addbook",AddBookHandler(dbops))
	router.GET("/getbookbytitle/:title",GetBookByTitleHandler(dbops))
	router.GET("/bookdetails",GetBookDetails(dbops))
        http.Handle("/",router)
	http.ListenAndServe(":8888",router)
}
