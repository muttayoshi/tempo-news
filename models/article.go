package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Article struct {
	Id        uint           `gorm:"primaryKey" json:"id"`
	Title     string         `gorm:"type:varchar(300)" json:"title"`
	HeadLine  string         `gorm:"type:varchar(1000)" json:"head_line"`
	Content   string         `gorm:"type:text" json:"content"`
	PostedAt  datatypes.Time `json:"posted_at" gorm:"autoCreateTime"`
	CreatedAt datatypes.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt datatypes.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
