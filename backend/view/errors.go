package view

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Unauthorized(c *gin.Context, obj any) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": obj,
	})
}
