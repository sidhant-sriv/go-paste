// main app using gin framework
package main

import (
	"fmt"
	_ "log"
	_ "net/http"
	_ "os"
	_ "time"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"

	"sidhant/pastebin-clone/db"
	"sidhant/pastebin-clone/handlers"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// setup main router
func main() {
	fmt.Println("Starting server...")
	db.InitDB()
	//setup router
	router := gin.Default()
	//setup routes
	router.GET("/", handlers.HelloHandler)
	router.POST("/", handlers.PostHandler)

	//paste routes
	router.POST("/paste", handlers.CreatePaste)
	router.GET("/pastes", handlers.GetAllPastes)
	router.GET("/paste/:id", handlers.GetPasteById)
	router.PUT("/paste/:id", handlers.UpdatePasteById)
	router.DELETE("/paste/:id", handlers.DeletePasteById)
	//run server
	router.Run(":8080")
}

//Curl requests to test paste routes
//curl -X POST -H "Content-Type: application/json" -d '{"title":"test","content":"test"}' http://localhost:8080/paste
//curl -X GET http://localhost:8080/pastes
//curl -X GET http://localhost:8080/paste/1

//Write SQL queries to make a bunch of pastes that look like actual pastes (make them joke title and joke content)
//INSERT INTO pastes (title, content) VALUES ('why did the chicken cross the road?', 'to get to the other side');
