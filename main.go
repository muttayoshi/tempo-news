package main

import (
	"github.com/gin-gonic/gin"
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

	//router.GET("/api/v1/articles", articleController.Index)

	err := router.Run(":8000")
	if err != nil {
		return
	}
}
