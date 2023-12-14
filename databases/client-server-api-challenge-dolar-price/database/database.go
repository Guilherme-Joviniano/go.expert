package database

import (
	"github.com/Guilherme-Joviniano/go-currency-api/data"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func Connect() (*Database, error) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(data.Currency{}, data.CurrencyCode{})

	return &Database{DB: db}, nil
}
