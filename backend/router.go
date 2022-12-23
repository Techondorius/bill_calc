package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	g := gin.Default()

	// Cors
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	g.Use(cors.New(config))

	g.GET("/", index)
	g.POST("/", index)

	return g
}

func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello",
	})
}
