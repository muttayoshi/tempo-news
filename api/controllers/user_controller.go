package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/muttayoshi/tempo-news/lib"
	"github.com/muttayoshi/tempo-news/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func SignUp(c *gin.Context) {
	//get email and pass from body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Something went wrong",
		})
		return
	}

	//hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to hash password",
		})
		return
	}

	//create user
	user := models.User{
		Email:    body.Email,
		Password: string(hash),
	}
	result := lib.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to hash password",
		})
		return
	}

	//response
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Sign Up success.",
	})
	//return
}
