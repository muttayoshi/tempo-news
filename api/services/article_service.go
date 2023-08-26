package services

import "github.com/muttayoshi/tempo-news/models"

type ArticleService interface {
	Create(article *models.Article) error
	GetAll([]*models.Article, error)
	GetDetail(*string) (*models.Article, error)
	Update(article *models.Article) error
	Delete(*string) error
}

type ArticleServiceImp struct {
}

func (a *ArticleServiceImp) Create(article *models.Article) error {
	return nil
}
