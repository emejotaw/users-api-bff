package config

import (
	"github.com/emejotaw/users-api-bff/internal/database/entity"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {

	dsn := "database.db"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&entity.Address{}, &entity.User{})

	if err != nil {
		panic(err)
	}

	return db
}
