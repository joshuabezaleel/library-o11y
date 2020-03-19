package book

import (
	"context"
	"errors"

	"github.com/joshuabezaleel/library-o11y/log"

	"github.com/sirupsen/logrus"
)

// Errors definition.
var (
	ErrGetAll  = errors.New("Failed retrieving all Books")
	ErrGetBook = errors.New("Failed retrieving particular Book")
)

// Service provides basic operations on Book domain model.
type Service interface {
	GetAll(ctx context.Context) ([]*Book, error)
	Get(ctx context.Context, bookID int) (*Book, error)
}

type service struct {
	bookRepository Repository

	logger *log.Logger
}

// NewBookService creates an instance of the service for the Book domain model
// with all of the necessary dependencies.
func NewBookService(bookRepository Repository, logger *log.Logger) Service {
	return &service{
		bookRepository: bookRepository,
		logger:         logger,
	}
}

func (s *service) GetAll(ctx context.Context) ([]*Book, error) {
	books, err := s.bookRepository.GetAll(ctx)
	if err != nil {
		logrus.Error(ErrGetAll)
		return nil, ErrGetAll
	}

	return books, nil
}

func (s *service) Get(ctx context.Context, bookID int) (*Book, error) {
	retrievedBook, err := s.bookRepository.Get(ctx, bookID)
	if err != nil {
		logrus.Errorf("%s ID: %v", ErrGetBook, bookID)
		return nil, ErrGetBook
	}

	return retrievedBook, nil
}
