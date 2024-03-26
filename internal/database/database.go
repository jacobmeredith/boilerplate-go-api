package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDatabase(path string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Example{})

	return db, nil
}
