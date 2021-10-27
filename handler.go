package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func test(c *gin.Context) {
	c.Set("aaa", 222222)
	c.Next()
}

func test2(c *gin.Context) {
	d, ok := c.Get("aaa")
	c.JSON(http.StatusOK, gin.H{
		"aaa":    d,
		"status": ok,
	})
}

func login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {

		response(c, "不正確參數")
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

func getUserInfo(c *gin.Context) {

	id, _, ok := getSession(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	user, err := findUserByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, user)
}
