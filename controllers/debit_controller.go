package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hyperyuri/webapi-with-go/datatase"
	"github.com/hyperyuri/webapi-with-go/models"
)

func ShowDebit(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be intenger",
		})

		return
	}

	db := datatase.GetDatabase()

	var debit models.DebitCard
	err = db.First(&debit, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find debit payment: " + err.Error(),
		})

		return
	}

	c.JSON(200, debit)
}
