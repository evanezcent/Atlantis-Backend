package main

import (
	"Atlantis-Backend/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting apps...")
	config.InitConnection()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
