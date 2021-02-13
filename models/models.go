package models

import "gorm.io/gorm"

// User ...
type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Notes[] Note `json:"notes" gorm:"foreignKey:User;references:Username"`
}

// Note ...
type Note struct {
	gorm.Model
	Title string `json:"title"`
	Text string `json:"text"`
	User string `json:"user"`
}

//LoginInfo ...
type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//NoteChange ...
type NoteChange struct {
	Title string `json:"title"`
	Text string `json:"text"`
}