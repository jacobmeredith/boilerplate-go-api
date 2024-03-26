package service

import (
	"github.com/jacobmeredith/boilerplate-go-api/internal/domain"
	"github.com/jacobmeredith/boilerplate-go-api/internal/store"
)

type ExampleService struct {
	store *store.ExampleStore
}

func NewExampleService(store *store.ExampleStore) *ExampleService {
	return &ExampleService{
		store: store,
	}
}

func (es *ExampleService) CreateExample(message string) error {
	domainExample, err := domain.NewExample(message)
	if err != nil {
		return err
	}

	err = es.store.CreateExample(domainExample.Message)
	if err != nil {
		return err
	}

	return nil
}

func (es *ExampleService) GetExample(id uint) (*domain.Example, error) {
	example, err := es.store.GetExample(id)
	if err != nil {
		return nil, err
	}

	domainExample, err := domain.NewExample(example.Message)
	if err != nil {
		return nil, err
	}

	return domainExample, nil
}
