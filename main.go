package main

import (
	"errors"
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct{
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Beyond Good and Evil", Author: "Fredrich Nietzhe", Quantity: 2},
	{ID: "2", Title: "Think and Grow Rich", Author: "Napoleon Hill", Quantity: 5},
	{ID: "3", Title: "The Changing World Order", Author: "Ray Dalio", Quantity: 6},
}

func bookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func getBookById(c *gin.Context) {
	id := c.Param("id")
	book, err := bookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
	}

	book, err := bookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Book not found."})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Not enough books to checkout."})		
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookById)
	router.POST("/books", createBook)
	router.PATCH("/checkout", checkoutBook)
	router.Run("localhost:8080")
}