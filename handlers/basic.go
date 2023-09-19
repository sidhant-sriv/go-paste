// Route handlers for just random json responses
package handlers

import (
	"fmt"
	"sidhant/pastebin-clone/models"

	gin "github.com/gin-gonic/gin"
)

// function that returns a json response
func HelloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}

// function that returns the same json body using the paste model on post request
func PostHandler(c *gin.Context) {
	var paste models.Paste
	err  := c.BindJSON(&paste)
	if err != nil {
		fmt.Println(err)
	}
	paste.ID = 100

	fmt.Println(paste)
	c.JSON(200, paste)
}

//curl request to test post handler
// curl -X POST -H "Content-Type: application/json" -d '{"id":1,"title":"test","content":"test"}' http://localhost:8080/
