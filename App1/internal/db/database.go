package db

import (
	"github.com/TalesPalma/GolangRabbitMQ/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

func LoadDatabase() {
	Db, err = gorm.Open(sqlite.Open("msgDatabase.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = Db.AutoMigrate(&models.Message{})

	if err != nil {
		panic("failed to migrate the database")
	}

}
