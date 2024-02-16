package controllers

import (
	"test_gin_go/initializers"
	"test_gin_go/models"

	"github.com/gin-gonic/gin"
)

func UsersCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Firstname string
		Lastname  string
		Email     string
		Password  string
	}
	c.Bind(&body)

	// Create a post
	user := models.User{Firstname: body.Firstname, Lastname: body.Lastname, Email: body.Email, Password: body.Password}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"user": user,
	})
}

func UsersIndex(c *gin.Context) {
	// Get the users
	var users []models.User
	initializers.DB.Find(&users)

	// Respond with them
	c.JSON(200, gin.H{
		"users": users,
	})
}

func UsersShow(c *gin.Context) {
	// Get id of url
	id := c.Param("id")

	// Get the user
	var user models.User
	initializers.DB.First(&user, id)

	// Return him
	c.JSON(200, gin.H{
		"user": user,
	})
}

func UsersUpdate(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Get the data off req body
	var body struct {
		Firstname string
		Lastname  string
		Email     string
		Password  string
	}
	c.Bind(body)

	// Find the user were updating
	var user models.User
	initializers.DB.First(&user, id)

	// Update it
	initializers.DB.Model(&user).Updates(models.User{Firstname: body.Firstname, Lastname: body.Lastname, Email: body.Email, Password: body.Password})

	// Return him
	c.JSON(200, gin.H{
		"user": user,
	})
}

func UsersDelete(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Delete the user
	initializers.DB.Delete(&models.User{}, id)

	// Respond
	c.Status(200)
}
