package main

import (
	"github.com/gin-gonic/gin"
	"github.com/muttayoshi/tempo-news/api/controllers"
	"github.com/muttayoshi/tempo-news/lib"
	"net/http"
)

func main() {
	router := gin.New()
	router.GET("/", func(context *gin.Context) {
		context.String(200, "Tempo News")
	})

	config := lib.NewConfig()
	lib.ConnectDatabase(config)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/api/v1/articles", controllers.Index)
	router.GET("/api/v1/article/:id", controllers.Show)
	router.POST("/api/v1/article", controllers.Create)
	router.PUT("/api/v1/article/:id", controllers.Update)
	router.DELETE("/api/v1/article", controllers.Delete)

	err := router.Run(":8000")
	if err != nil {
		return
	}
}
