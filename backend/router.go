package main

import (
	"math/rand"
	"time"

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
	g.GET("/rand", randRouter)

	return g
}

func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello",
	})
}

func randRouter(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	c.JSON(200, gin.H{
		"message": rand.Int(),
	})
}
