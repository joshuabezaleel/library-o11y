package book

import (
	"errors"
)

// Errors definition.
var (
	ErrGetAll = errors.New("Failed retrieving all Books")
)

// Service provides basic operations on Book domain model.
type Service interface {
	GetAll() ([]*Book, error)
}

type service struct {
	bookRepository Repository
}

// NewBookService creates an instance of the service for the Book domain model
// with all of the necessary dependencies.
func NewBookService(bookRepository Repository) Service {
	return &service{
		bookRepository: bookRepository,
	}
}

func (s *service) GetAll() ([]*Book, error) {
	books, err := s.bookRepository.GetAll()
	if err != nil {
		return nil, ErrGetAll
	}

	return books, nil
}
