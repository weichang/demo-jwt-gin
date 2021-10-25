package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main()  {

	db,err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}
	DB = db
	//_, _ = insertUser("Jeffrey", "12345", 30)
	//_, _ = insertUser("David", "12345", 25)

}