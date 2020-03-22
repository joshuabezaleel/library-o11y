package persistence

import (
	"context"
	"errors"
	"fmt"

	"github.com/joshuabezaleel/library-o11y/book"
	"github.com/joshuabezaleel/library-o11y/log"

	"github.com/fluent/fluent-logger-golang/fluent"
)

var allBooks = []*book.Book{
	&book.Book{ID: 0, Title: "Medic Without Shame", Author: "Richard Gizmo Seventh"},
	&book.Book{ID: 1, Title: "Blacksmith Of The Forest", Author: "Mark R. Steel"},
	&book.Book{ID: 2, Title: "Men With Determination", Author: "Maxwell Southwell"},
	&book.Book{ID: 3, Title: "Snakes Of My Imagination", Author: "M. R. Smitherd"},
	&book.Book{ID: 4, Title: "Soldiers And Witches", Author: "Gizmo Richards"},
	&book.Book{ID: 5, Title: "Lions And Men", Author: "Richard Gizmo Seventh"},
	&book.Book{ID: 6, Title: "Restoration Of Wood", Author: "Marissa Smirnova"},
	&book.Book{ID: 7, Title: "Curse Of The River", Author: "Maree Markson"},
	&book.Book{ID: 8, Title: "Write About Technology", Author: "Smith Markblood"},
	&book.Book{ID: 9, Title: "Duke Of The North", Author: "Mathias R. Seventh"},
	&book.Book{ID: 10, Title: "Tree Of Darkness", Author: "Ocean Roxie Northern"},
	&book.Book{ID: 11, Title: "Humans Of The Ancestors", Author: "Wendell Billerbeck"},
	&book.Book{ID: 12, Title: "Fish Of The Gods", Author: "Mathias R. Seventh"},
	&book.Book{ID: 13, Title: "Doctors And Owls", Author: "Wearmouth O. Northern"},
	&book.Book{ID: 14, Title: "Hunters And Assassins", Author: "Octavia Birszwilks"},
	&book.Book{ID: 15, Title: "Surprise With Honor", Author: "Richard Gizmo Seventh"},
	&book.Book{ID: 16, Title: "Choice Of Yesterday", Author: "Bishop Wenblood"},
	&book.Book{ID: 17, Title: "Sounds In The East", Author: "Wes O. Armedrobber"},
	&book.Book{ID: 18, Title: "Crying In The Stars", Author: "M. R. Smitherd"},
	&book.Book{ID: 19, Title: "King With Sins", Author: "Joshua Bezaleel Abednego"},
}

type bookRepository struct {
	logger *log.Logger

	fluentLogger *fluent.Fluent
}

// NewBookRepository returns initialized implementations of the repository for
// Book domain model.
func NewBookRepository(logger *log.Logger, fluentLogger *fluent.Fluent) book.Repository {
	return &bookRepository{
		logger:       logger,
		fluentLogger: fluentLogger,
	}
}

func (repo *bookRepository) GetAll(ctx context.Context) ([]*book.Book, error) {
	var books []*book.Book

	for _, eachBook := range allBooks {
		books = append(books, eachBook)
	}

	repo.logger.Log.Debug("Query GetAll")
	repo.fluentLogger.Post("repository", "Query GetAll")

	return books, nil
}

func (repo *bookRepository) Get(ctx context.Context, bookID int) (*book.Book, error) {
	if bookID > len(allBooks)-1 || allBooks[bookID] == nil {
		repo.logger.Log.Errorf("Error: Query Get ID %v", bookID)
		return nil, errors.New("Book not found")
	}

	logMessage := fmt.Sprintf("Query Get ID %v", bookID)
	repo.logger.Log.Debugf(logMessage)
	repo.fluentLogger.Post("repository", logMessage)

	return allBooks[bookID], nil
}
