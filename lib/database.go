package lib

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Database struct {
	ORM *gorm.DB
}

var DB *gorm.DB

func ConnectDatabase(config Config) Database {
	dbConfig := postgres.Config{
		DSN: config.Database.DSN(),
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
	if err != nil {
		panic(err)
	}

	//errMigrate := db.AutoMigrate(
	//	&models.Article{},
	//	&models.User{},
	//)
	//if errMigrate != nil {
	//	panic(errMigrate)
	//}

	DB = db

	return Database{
		ORM: db,
	}
}
