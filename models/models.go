//Paste post model
package models
type Paste struct {
	//make id primary key and autoincrement
	ID      int    `gorm:"primary_key;auto_increment" json:"id"`
	Title  string `json:"title"`
	Content string `json:"content"`
}

//User model
type User struct {
	//make id primary key and autoincrement
	ID      int    `gorm:"primary_key;auto_increment" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Pastes  []Paste `json:"pastes"`
}
