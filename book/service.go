package book

import (
	"errors"
	"sort"
)

// Errors definition.
var (
	ErrGetAll          = errors.New("Failed retrieving all Books")
	ErrGetBookTopicIDs = errors.New("Failed retrieving Book's topic IDs")
	ErrGetTopicsByID   = errors.New("Failed retrieving topics")
)

// Service provides basic operations on Book domain model.
type Service interface {
	GetAll(topicIDs []int, sort string, order string) ([]*Book, error)
	GetBookTopicIDs(bookID int) ([]int, error)
	GetTopicsByID(topicIDs []int) ([]string, error)
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

func (s *service) GetAll(topicIDs []int, sortParam string, order string) ([]*Book, error) {
	books, err := s.bookRepository.GetAll(topicIDs)
	if err != nil {
		return nil, ErrGetAll
	}

	if len(topicIDs) != 4 {
		for _, book := range books {
			bookTopicIDs, err := s.GetBookTopicIDs(book.ID)
			if err != nil {
				return nil, ErrGetBookTopicIDs
			}

			bookTopics, err := s.GetTopicsByID(bookTopicIDs)
			if err != nil {
				return nil, ErrGetTopicsByID
			}

			book.Topics = bookTopics
		}
	}

	if sortParam == "Title" {
		if order == "asc" {
			sort.Sort(ByTitle(books))
		} else if order == "desc" {
			sort.Sort(sort.Reverse(ByTitle(books)))
		}
	} else if sortParam == "Author" {
		if order == "asc" {
			sort.Sort(ByAuthor(books))
		} else if order == "desc" {
			sort.Sort(sort.Reverse(ByAuthor(books)))
		}
	}

	return books, nil
}

func (s *service) GetBookTopicIDs(bookID int) ([]int, error) {
	bookTopicIDs, err := s.bookRepository.GetBookTopicIDs(bookID)
	if err != nil {
		return nil, ErrGetBookTopicIDs
	}

	return bookTopicIDs, nil
}

func (s *service) GetTopicsByID(topicIDs []int) ([]string, error) {
	topics, err := s.bookRepository.GetTopicsByID(topicIDs)
	if err != nil {
		return nil, ErrGetTopicsByID
	}

	return topics, nil
}

// ByTitle implements sort.Interface based on the Title field.
type ByTitle []*Book

func (books ByTitle) Len() int           { return len(books) }
func (books ByTitle) Less(i, j int) bool { return books[i].Title < books[j].Title }
func (books ByTitle) Swap(i, j int)      { books[i], books[j] = books[j], books[i] }

// ByAuthor implements sort.Interface based on the Author field.
type ByAuthor []*Book

func (books ByAuthor) Len() int           { return len(books) }
func (books ByAuthor) Less(i, j int) bool { return books[i].Author < books[j].Author }
func (books ByAuthor) Swap(i, j int)      { books[i], books[j] = books[j], books[i] }
