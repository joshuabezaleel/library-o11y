package main

import (
	"github.com/joshuabezaleel/library-o11y/book"
	"github.com/joshuabezaleel/library-o11y/log"
	"github.com/joshuabezaleel/library-o11y/persistence"
	"github.com/joshuabezaleel/library-o11y/server"
)

func main() {
	logger := log.NewLogger()

	bookRepository := persistence.NewBookRepository(logger)

	bookService := book.NewBookService(bookRepository, logger)

	srv := server.NewServer(bookService, logger)
	srv.Run("8082")
}
