package view

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StatusOK(c *gin.Context, obj any) {
	c.JSON(http.StatusOK, obj)
}
