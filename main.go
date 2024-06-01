package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/volley", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "volley",
		})
	})

	r.GET("/box", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "box",
		})
	})

	log.Fatalln(r.Run(":8080"))
}
