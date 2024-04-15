package main

import (
	// "errors"
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

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.Run("localhost:8080")
}