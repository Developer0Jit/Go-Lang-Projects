package main

import (
	"crud_api/initializers"
	"crud_api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	initializers.DB.AutoMigrate(models.Post{})
}
