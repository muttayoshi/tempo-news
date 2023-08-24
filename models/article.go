package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Id        uint           `gorm:"primaryKey" json:"id"`
	Title     string         `gorm:"type:varchar(300)" json:"title"`
	HeadLine  string         `gorm:"type:varchar(1000)" json:"head_line"`
	Content   string         `gorm:"type:text" json:"content"`
	PostedAt  datatypes.Time `json:"posted_at"`
	CreatedAt datatypes.Time `json:"created_at"`
	UpdatedAt datatypes.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
