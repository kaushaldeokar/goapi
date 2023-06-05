package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// Struct tags such as json:"artist" specify what a field’s name should be when the struct’s contents are serialized into JSON. Without them, the JSON would use the struct’s capitalized field names – a style not as common in JSON.
type Book struct {
	ID     string
	Title  string
	Author string
	Price  int
}

var Library = []Book{
	{ID: "1", Title: "Harry Potter ", Author: "J.K Rowling", Price: 200},
	{ID: "2", Title: "Game of Thrones ", Author: "George R Martin", Price: 1500},
}

func getBooks(c *gin.Context) {

	// Context is the most important part of gin.
	// It allows us to pass variables between middleware,
	// manage the flow, validate the JSON of a request and
	// render a JSON response for example.
	c.IndentedJSON(http.StatusOK, Library)

}
func addBooks(c *gin.Context) {
	var newBook Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	
	Library = append(Library, newBook)
	// c.IndentedJSON(http.StatusCreated, newBook)
	c.IndentedJSON(http.StatusCreated, "recieved")

}

func getBooksbyID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range Library {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }

	c.IndentedJSON(http.StatusOK, "not found try again")

}
func main() {
	router := gin.Default()

	router.GET("/Books", getBooks)
	router.GET("/Books/:id", getBooksbyID)
	router.POST("/add", addBooks)
	router.Run("localhost:8000")
}
