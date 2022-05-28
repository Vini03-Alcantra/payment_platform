package controllers

import "github.com/gin-gonic/gin"

func ShowDebit(c *gin.Context) {
	c.JSON(200, gin.H{
		"value": "ok",
	})
}
