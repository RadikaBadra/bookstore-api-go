package models

import (
	"gorm.io/gorm"
)

type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  int    `json:"price"`
}

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
}