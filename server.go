package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello World!"})
	})
	r.Run()
}
