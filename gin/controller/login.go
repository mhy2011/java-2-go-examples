package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginJson(c *gin.Context) {
	var json Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if json.Username != "root" || json.Password != "root" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "invalid username or password"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK"})

}

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}