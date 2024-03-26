package domain

import "errors"

type ExampleMessage string

type Example struct {
	Message ExampleMessage
}

func NewExampleMessage(message string) (ExampleMessage, error) {
	if len(message) < 1 {
		return "", errors.New("message length must be greater than 0")
	}

	return ExampleMessage(message), nil
}

func NewExample(message string) (*Example, error) {
	domainMessage, err := NewExampleMessage(message)
	if err != nil {
		return nil, err
	}

	return &Example{
		Message: domainMessage,
	}, nil
}
