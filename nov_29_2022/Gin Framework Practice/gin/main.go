package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func HandlFunc(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello gin",
	})
}

func main() {
	r := gin.Default()
	r.GET("/home", HandlFunc)
	r.GET("/home/foo", func(c *gin.Context) {
		c.JSON(200, "foo")
	})
	log.Fatal(r.Run(":9000"))
}
