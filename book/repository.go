package book

// Repository provides access to the Book store.
type Repository interface {
	GetAll(topicIDs []int) ([]*Book, error)
	GetBookTopicIDs(bookID int) ([]int, error)
	GetTopicsByID(topicIDs []int) ([]string, error)
}
