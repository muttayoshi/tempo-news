package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/muttayoshi/tempo-news/api/controllers"
	"github.com/muttayoshi/tempo-news/api/middlewares"
	"github.com/muttayoshi/tempo-news/lib"
	"log"
	"os"
)

func init() {
	lib.ConnectDatabase(lib.NewConfig())
	lib.MigrateDatabase()
}

func main() {
	errGetEnv := godotenv.Load()
	if errGetEnv != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.New()
	router.GET("/", func(context *gin.Context) {
		context.String(200, "Tempo News")
	})

	router.POST("/api/v1/signup", controllers.SignUp)
	router.POST("/api/v1/login", controllers.SignIn)
	router.GET("/api/v1/validate", middlewares.Auth, controllers.Validate)
	router.GET("/api/v1/articles", middlewares.Auth, controllers.Index)
	router.GET("/api/v1/article/:id", middlewares.Auth, controllers.Show)
	router.POST("/api/v1/article", middlewares.Auth, controllers.Create)
	router.PUT("/api/v1/article/:id", middlewares.Auth, controllers.Update)
	router.DELETE("/api/v1/article", middlewares.Auth, controllers.Delete)

	err := router.Run(":" + os.Getenv("PORT"))
	if err != nil {
		return
	}
}
