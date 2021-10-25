package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {

		response (c,"不正確參數")
		return
		//c.JSON(http.StatusBadRequest, gin.H{
		//	"error": "不正確參數",
		//})
		//return
	}

	user, err := findUserByUsername(req.Username)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "找不到帳號",
		})
		return
	}
	if user.Password != req.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "密碼不正確",
		})
		return
	}

	token, err := generateToken(*user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
