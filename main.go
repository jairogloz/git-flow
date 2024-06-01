package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	r.GET("/soccer", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "soccer",
		})
	})

	log.Fatalln(r.Run(":8080"))
}
