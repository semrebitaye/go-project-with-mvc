package main

import (
	"mvc-go/initializers"
	"mvc-go/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
