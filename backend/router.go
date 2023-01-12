package main

import (
	"math/rand"
	"time"

	"github.com/Techondorius/bill_calc/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	g := gin.Default()

	// Cors
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		MaxAge:           300,
	}
	config.AddAllowHeaders(
		"Access-Control-Allow-Credentials",
		"Access-Control-Allow-Headers",
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"Authorization",
	)
	g.Use(cors.New(config))

	g.GET("/", index)
	g.POST("/", index)
	g.GET("/rand", randRouter)

	g.POST("/login", controller.Login)

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
