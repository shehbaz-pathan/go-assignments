package main 
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"log"
	"fmt"
	"strings"
)
func GetBooksHandler(ops DBOperations) gin.HandlerFunc {
	return func(c *gin.Context) {
		var books = []Book{}
		books,err:=ops.GetBooks()
		if err!=nil {
                        log.Printf("%s while reading data",err)
                        return
                }
                c.IndentedJSON(http.StatusOK,books)

	}
}

func AddBookHandler(ops DBOperations) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book Book
		if err:= c.ShouldBindJSON(&book);err!=nil {
			log.Printf("%s while binding",err)
			c.IndentedJSON(http.StatusBadRequest,gin.H{"error":"please check the input data below are the require fields","fields":"isbn(number), title(text), synopsis(text), authorname(text), price(number)"})
			return
		}
		err:= ops.AddBook(book)
                if err != nil {
                        log.Printf("%s while adding book",err)
                        if strings.Contains(strings.ToLower(fmt.Sprintf("%s",err)),"duplicate") {
				c.JSON(http.StatusBadRequest,gin.H{"message": "duplicate book"})
                                return
                        }
                        return

                }
                c.IndentedJSON(http.StatusCreated,book)

	}
}

func GetBookByTitleHandler(ops DBOperations) gin.HandlerFunc {
	return func(c *gin.Context) {
		var books []Book
		title:=c.Param("title")
		books,err:=ops.GetBookByTitle(title)
		if err!= nil {
			log.Printf("%s while getting book details",err)
			c.JSON(http.StatusInternalServerError,gin.H{"message": err})
			return
		}
		if len(books) <=0 {
			c.JSON(http.StatusNotFound,gin.H{"message": fmt.Sprintf("404 book by title %s not found",title)})
			return
		}
		c.IndentedJSON(http.StatusOK,books)
	}
}
func GetBookDetails(ops DBOperations) gin.HandlerFunc {
	return func(c *gin.Context) {
                var books = []Book{}
                books,err:=ops.GetBooks()
                if err!=nil {
                        log.Printf("%s while reading data",err)
                        return
                }
                c.HTML(http.StatusOK,"index.html",books)

        }
}
