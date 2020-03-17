package persistence

import (
	"errors"

	"github.com/joshuabezaleel/library-o11y/book"
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

type bookRepository struct{}

// NewBookRepository returns initialized implementations of the repository for
// Book domain model.
func NewBookRepository() book.Repository {
	return &bookRepository{}
}

func (repo *bookRepository) GetAll() ([]*book.Book, error) {
	var books []*book.Book

	for _, eachBook := range allBooks {
		books = append(books, eachBook)
	}

	return books, nil
}

func (repo *bookRepository) Get(bookID int) (*book.Book, error) {
	if bookID > len(allBooks)-1 {
		return nil, errors.New("Book not found")
	}

	return allBooks[bookID], nil
}
