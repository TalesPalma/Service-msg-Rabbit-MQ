package db

import (
	"github.com/TalesPalma/App2/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

func InitDatabase() {
	Db, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	if err = Db.AutoMigrate(&models.Message{}); err != nil {
		panic("failed to migrate database")
	}

}
