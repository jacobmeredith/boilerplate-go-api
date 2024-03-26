package store

import (
	"github.com/jacobmeredith/boilerplate-go-api/internal/database"
	"github.com/jacobmeredith/boilerplate-go-api/internal/domain"
	"gorm.io/gorm"
)

type ExampleStore struct {
	db *gorm.DB
}

func NewExampleStore(db *gorm.DB) *ExampleStore {
	return &ExampleStore{
		db: db,
	}
}

func (es *ExampleStore) CreateExample(message domain.ExampleMessage) error {
	result := es.db.Create(&database.Example{Message: string(message)})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (es *ExampleStore) GetExample(id uint) (*database.Example, error) {
	example := new(database.Example)

	result := es.db.First(&example, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return example, nil
}
