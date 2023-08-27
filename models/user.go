package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string         `gorm:"column:username;size:64;not null;index;unique" json:"username" validate:"required"`
	Password  string         `gorm:"column:password;not null;" json:"password" validate:"required"`
	Email     string         `gorm:"column:email;not null;unique;" json:"email" validate:"required"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type UserIndexResult struct {
	Results []User `json:"list"`
}
