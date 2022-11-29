package main

import (
	"crud_api/controllers"
	"crud_api/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.UpdatePosts)
	r.GET("/posts", controllers.GetAllPosts)
	r.GET("/posts/:id", controllers.GetSinglePost)
	r.DELETE("/posts/:id", controllers.DeletePosts)

	r.Run() // listen and serve on 0.0.0.0:8080
}
