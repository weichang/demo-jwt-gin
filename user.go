package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" `
	Password string `json:"password"`
	Age      uint   `json:"age"`
}

type FindUser struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Age      uint   `json:"age"`
}
