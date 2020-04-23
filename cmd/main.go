package main

import (
	"github.com/joshuabezaleel/library-o11y/book"
	"github.com/joshuabezaleel/library-o11y/log"
	"github.com/joshuabezaleel/library-o11y/persistence"
	"github.com/joshuabezaleel/library-o11y/server"

	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/sirupsen/logrus"
)

func main() {
	// time.Sleep(5 * time.Second)
	logger := log.NewLogger()
	logger.Log.SetFormatter(&logrus.JSONFormatter{})

	fluentLogger, err := fluent.New(fluent.Config{FluentPort: 24224, FluentHost: "fluentd"})
	if err != nil {
		logger.Log.Fatal(err)
	}
	defer fluentLogger.Close()

	bookRepository := persistence.NewBookRepository(logger, fluentLogger)

	bookService := book.NewBookService(bookRepository, logger, fluentLogger)

	srv := server.NewServer(bookService, logger, fluentLogger)
	srv.Run("8082")
}
