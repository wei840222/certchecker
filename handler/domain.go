package handler

import (
	"github.com/wei840222/certchecker/db"

	"errors"
	"net/http"
	"strconv"

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

func DeleteDomain(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if id <= 0 {
		err := errors.New("id should be natual number")
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DeleteDomain(uint(id)); err != nil {
		panic(err)
	}

	c.Status(http.StatusNoContent)
}
