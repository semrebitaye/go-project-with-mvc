package controllers

import (
	"log"
	"mvc-go/initializers"
	"mvc-go/models"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	var users []models.User

	initializers.DB.Find(&users)
	c.IndentedJSON(200, gin.H{
		"User": users,
	})
}

func CreateUser(c *gin.Context) {
	var user models.User
	var body struct {
		Name    string
		Email   string
		Age     int
		Address string
	}
	c.Bind(&body)

	user = models.User{Name: body.Name, Email: body.Email, Age: body.Age, Address: body.Address}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		log.Fatal("Failed to create user")
	}
	c.IndentedJSON(200, gin.H{
		"User": user,
	})
}

func GetUserById(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	initializers.DB.First(&user, id)

	c.IndentedJSON(200, gin.H{
		"User": user,
	})
}

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	initializers.DB.Delete(&user, id)

	c.IndentedJSON(200, gin.H{
		"User": user,
	})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	var body struct {
		Name    string
		Email   string
		Age     int
		Address string
	}
	c.Bind(&user)
	id := c.Param("id")
	initializers.DB.First(&user, id)

	initializers.DB.Model(&user).Updates(models.User{Name: body.Name, Email: body.Email, Age: body.Age, Address: body.Address})

	c.IndentedJSON(200, gin.H{
		"User": user,
	})
}
