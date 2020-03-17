package book

// Book domain model.
type Book struct {
	ID     int      `json:"id"`
	Title  string   `json:"title"`
	Author string   `json:"author"`
	Topics []string `json:"topics"`
}

// NewBook creates a new instance of Book domain model.
func NewBook(id int, title string, author string, topics []string) *Book {
	return &Book{
		ID:     id,
		Title:  title,
		Author: author,
		Topics: topics,
	}
}
