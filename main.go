package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

var DB *gorm.DB

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}
	DB = db
	//_, _ = insertUser("Jeffrey", "12345", 30)
	//_, _ = insertUser("David", "12345", 25)

	r := gin.Default()
	r.POST("/login", login)

	r.Use(test)
	r.GET("/test2", test2)


	r.Use(verifyToken)
	r.GET("/info", getUserInfo)

	if err := r.Run("localhost:5688"); err != nil {
		panic(err)
	}
}
func verifyToken(c *gin.Context) {

	token, ok := getToken(c)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}

	id, username, err := validateToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}

	c.Set("id", id)
	c.Set("username", username)
	c.Next()

}

func getToken(c *gin.Context) (string, bool) {
	authValue := c.GetHeader("Authorization")

	arr := strings.Split(authValue, " ")
	if len(arr) != 2 {
		return "", false
	}

	authType := strings.Trim(arr[0], "\n\r\t")
	if strings.ToLower(authType) != strings.ToLower("Bearer") {
		return "", false
	}

	return strings.Trim(arr[1], "\n\t\r"), true

}

func getSession(c *gin.Context) (uint, string, bool) {
	id, ok := c.Get("id")
	if !ok {
		return 0, "", false
	}
	username, ok := c.Get("username")
	if !ok {
		return 0, "", false
	}
	return id.(uint), username.(string), true
}
