package lib

import (
	"github.com/muttayoshi/tempo-news/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Database struct {
	ORM *gorm.DB
}

func ConnectDatabase(config Config) Database {
	dsn := "host=localhost user=tempo_user password=tempo_password dbname=tempo_db port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	dbConfig := postgres.Config{
		DSN: dsn,
	}

	db, err := gorm.Open(postgres.New(dbConfig), &gorm.Config{
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   config.Database.TablePrefix + "_",
		},
		QueryFields: true,
	})

	db.AutoMigrate(&models.Article{})

	if err != nil {
		panic(err)
	}

	return Database{
		ORM: db,
	}
}
