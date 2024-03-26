package database

import "gorm.io/gorm"

type Example struct {
	gorm.Model
	Message string
}
