package lib

import "github.com/muttayoshi/tempo-news/models"

func MigrateDatabase() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Article{},
	)
	if err != nil {
		panic(err)
	}
}
