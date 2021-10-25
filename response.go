package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Data struct {
	Message  string `json:"message"`
	Message2 string `json:"message2"`
}

func response(c *gin.Context, message string) (res string) {
	c.JSON(http.StatusOK, gin.H{
		"serverTime":    time.Now().Format("2006-01-02 15:04:05"),
		"statusCode":    "0000",
		"statusMessage": message,
		"data": Data{
			Message: message,
			Message2: message,
		}})
	return
}
