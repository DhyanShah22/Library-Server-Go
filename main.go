package main

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// A gin router to handle requests

	var router *gin.Engine = gin.Default()

	router.GET("/books", getBooks)
	router.POST("/book", addBook)
	router.GET("/book/:id", getSingleBook)
	router.DELETE("/book/:id", deleteBook)
	router.PUT("/book/:id", updateBook)

	router.Run(":7000")
}

type book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
}

var library = []book{
	{
		ID: "1",
		Title: "System Design",
		Author: "Lee, Andrew S.",
	},
	{
		ID: "2",
		Title: "Introduction to Algorithms",
		Author: "Cormen",
	},
	{
		ID: "3",
		Title: "Introduction to DMS",
		Author: "John Doe",
	},
}

func getBooks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, library)
}

func addBook(context *gin.Context) {
	var newBook book

	if err:= context.BindJSON(&newBook); err!= nil{
		fmt.Println(err)

		context.IndentedJSON(http.StatusBadRequest, gin.H{
			"message" : "Invalid Request. Please provide a valid request body." ,
		})

		return
	}

	library = append(library, newBook)
	context.IndentedJSON(http.StatusCreated, newBook)
}

func getSingleBook(context *gin.Context) {
	var id string = context.Param("id")

	for _,p := range library {
		if p.ID == id {
			context.IndentedJSON(http.StatusOK, p)
			return
		}
	}

	context.IndentedJSON(
		http.StatusNotFound,
		gin.H{
			"message": "No Book Found with the given ID",
			},
	)
}

func deleteBook(context *gin.Context) {
	id := context.Param("id")

	for i, b := range library {
		if b.ID == id {
			library = append(library[:i], library[i+1:]...)
			context.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
			return
		}
	}

	context.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func updateBook(context *gin.Context) {
	id := context.Param("id")
	var updatedBook book

	if err := context.BindJSON(&updatedBook); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request. Please provide a valid request body."})
		return
	}

	for i, b := range library {
		if b.ID == id {
			library[i] = updatedBook
			context.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
			return
		}
	}

	context.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}