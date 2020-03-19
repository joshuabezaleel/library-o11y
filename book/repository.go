package book

import "context"

// Repository provides access to the Book store.
type Repository interface {
	GetAll(ctx context.Context) ([]*Book, error)
	Get(ctx context.Context, bookID int) (*Book, error)
}
