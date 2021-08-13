package main

import (
	"belajar-be-dasar-gin/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong Get",
		})
	})
	r.POST("/ping", func(c *gin.Context) {
		name := c.Query("name")
		c.JSON(200, gin.H{
			"name":    name,
			"message": "pong POST",
		})
	})
	r.GET("/user", controller.GetUser)
	r.POST("/user", controller.PostUser)
	r.Run(":9000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
