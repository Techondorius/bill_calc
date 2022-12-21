package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	// Cors
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	// config.AllowOrigins = []string{"http://0.0.0.0:3000", "http://localhost:3000"}
	g.Use(cors.New(config))
	g.GET("/", index)
	g.POST("/", index)
	g.Run()
	fmt.Println("Hello")
}

func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello",
	})
}
