package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/muttayoshi/tempo-news/lib"
	"github.com/muttayoshi/tempo-news/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func SignUp(c *gin.Context) {
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

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to hash password",
		})
		return
	}

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

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Sign Up success.",
	})
}

func SignIn(c *gin.Context) {
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

	var user models.User
	lib.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid email or password",
		})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid email or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, errToken := token.SignedString([]byte(os.Getenv("SECRET")))

	if errToken != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Something went wrong.",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Success",
		"token":   tokenString,
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Success.",
		"user":    user,
	})
}
