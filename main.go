package main

import (
	"mvc-go/controllers"
	"mvc-go/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()
	r.GET("/get", controllers.GetUser)
	r.POST("/create", controllers.CreateUser)
	r.GET("/get/:id", controllers.GetUserById)
	r.DELETE("/delete/:id", controllers.DeleteUser)
	r.PUT("/update/:id", controllers.UpdateUser)
	r.Run()
}
