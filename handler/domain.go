package handler

import (
	"github.com/wei840222/certchecker/db"

	"net/http"

	"github.com/gin-gonic/gin"
)

func ListDomain(c *gin.Context) {
	ds, err := db.ListDomain()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, ds)
}

func CreateDomain(c *gin.Context) {
	var d db.Domain
	if err := c.ShouldBindJSON(&d); err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.CreateDomain(&d); err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, d)
}
