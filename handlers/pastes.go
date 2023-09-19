// handlers for pastes
// Path: handlers/pastes.go
package handlers

import (
	"fmt"
	"sidhant/pastebin-clone/models"

	//gorm

	//postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"sidhant/pastebin-clone/db"

	gin "github.com/gin-gonic/gin"
)

// function that creates a paste
func CreatePaste(c *gin.Context) {

	//setup db
	db, err := db.GetDB()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	//create paste
	var paste models.Paste
	err = c.BindJSON(&paste)
	if err != nil {
		fmt.Println(err)
	}
	db.Create(&paste)
	//return json response
	c.JSON(200, paste)
}

// function that returns all pastes
func GetAllPastes(c *gin.Context) {
	//setup db
	db, err := db.GetDB()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	//get all pastes
	var pastes []models.Paste
	db.Find(&pastes)
	//return json response
	c.JSON(200, pastes)
}

// function that returns a paste by id
func GetPasteById(c *gin.Context) {
	//setup db
	db, err := db.GetDB()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	//get paste by id
	var paste models.Paste
	id := c.Param("id")
	db.First(&paste, id)
	//return json response
	c.JSON(200, paste)
}

// function that updates a paste by id
func UpdatePasteById(c *gin.Context) {
	//setup db
	db, err := db.GetDB()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	//update paste by id
	var paste models.Paste
	id := c.Param("id")
	db.First(&paste, id)
	err = c.BindJSON(&paste)
	if err != nil {
		fmt.Println(err)
	}
	db.Save(&paste)
	//return json response
	c.JSON(200, paste)
}

// function that deletes a paste by id and takes db as an argument
func DeletePasteById(c *gin.Context) {
	//setup db
	db, err := db.GetDB()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	//delete paste by id
	var paste models.Paste
	id := c.Param("id")
	db.First(&paste, id)
	db.Delete(&paste)
	//return json response
	c.JSON(200, paste)
}
