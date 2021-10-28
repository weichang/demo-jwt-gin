package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/weichang/demo-jwt-gin/docs"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"strings"
	_ "github.com/joho/godotenv/autoload"
)

var DB *gorm.DB

// @title demo-jwt-gin swagger
// @version 1.0
// @description demo-jwt-gin swagger

// @contact.name Jeffrey
// @contact.url

// @host localhost:5688
// schemes http

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
	r.Use(cors.Default())
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		// user login
		v1.POST("/login", Login)

		// get user info
		v1.Use(verifyToken)
		v1.GET("/info", getUserInfo)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
