package controllers

import (
	"crud_api/initializers"
	"crud_api/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	//Get data of request body

	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//Return in

	c.JSON(200, gin.H{
		"post": post,
	})

}

func GetAllPosts(c *gin.Context) {
	// Get all posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//Respond with team
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetSinglePost(c *gin.Context) {
	//get id of url
	id := c.Param("id")

	//Get single data
	var post models.Post
	initializers.DB.First(&post, id)

	//Respond with team
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func UpdatePosts(c *gin.Context) {
	//Get the id of url
	id := c.Param("id")

	//Get the data of request body

	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	//Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	//Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	//Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePosts(c *gin.Context) {
	//Get the id of url
	id := c.Param("id")

	//delete the post
	initializers.DB.Delete(&models.Post{}, id)

	c.Status(200)
}
