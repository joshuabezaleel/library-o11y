package main

import (
	"github.com/joshuabezaleel/library-o11y/book"
	"github.com/joshuabezaleel/library-o11y/persistence"
	"github.com/joshuabezaleel/library-o11y/server"
)

func main() {
	// var log = logrus.New()

	bookRepository := persistence.NewBookRepository()

	bookService := book.NewBookService(bookRepository)

	srv := server.NewServer(bookService)
	srv.Run("8082")
}
