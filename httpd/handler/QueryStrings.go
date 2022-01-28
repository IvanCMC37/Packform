package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func QueryStrings() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Query("name")
		product := c.Query("product")
		c.JSON(http.StatusOK, gin.H{
			"name":    name,
			"product": product,
		})
	}
}
