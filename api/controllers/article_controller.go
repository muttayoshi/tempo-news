package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/muttayoshi/tempo-news/lib"
	"github.com/muttayoshi/tempo-news/models"
	"gorm.io/gorm"
	"net/http"
)

type ArticleController struct {
	db lib.Database
}

func Index(c *gin.Context) {
	var articles []models.Article
	lib.DB.Find(&articles)
	c.JSON(http.StatusOK, gin.H{
		"data":    articles,
		"success": true,
		"message": "Success.",
	})
}

func Show(c *gin.Context) {
	var article models.Article
	id := c.Param("id")

	if err := lib.DB.First(&article, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"data":    nil,
				"success": false,
				"message": "Data not found.",
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"data":    nil,
				"success": false,
				"message": "Something went wrong.",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    article,
		"message": "Success.",
		"success": true,
	})
}

func Create(c *gin.Context) {
	var article models.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Something went wrong.",
		})
		return
	}

	lib.DB.Create(&article)
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Success.",
		"data":    article,
	})
}

func Update(c *gin.Context) {
	var article models.Article
	id := c.Param("id")
	if err := c.ShouldBindJSON(&article); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Something went wrong.",
		})
		return
	}

	if lib.DB.Model(&article).Where("id = ?", id).Updates(&article).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Cant update article.",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Success.",
		"data":    &article,
	})
}

func Delete(c *gin.Context) {
	var article models.Article

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Something went wrong.",
		})
		return
	}
	_id, _ := input.Id.Int64()
	if lib.DB.Delete(&article, _id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Cant delete data.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"messsage": "Success",
		"data":     article,
	})
}
