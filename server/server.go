package server

import (
	"context"
	"net/http"

	"github.com/joshuabezaleel/library-o11y/book"
	"github.com/joshuabezaleel/library-o11y/log"
	opentracing "github.com/opentracing/opentracing-go"

	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/gorilla/mux"
)

// Server holds dependencies for the Book HTTP server.
type Server struct {
	ctx    context.Context
	tracer opentracing.Tracer

	bookService book.Service

	Router *mux.Router

	Logger *log.Logger

	FluentLogger *fluent.Fluent
}

// NewServer returns a new Book HTTP server
// with all of the necessary dependencies.
func NewServer(ctx context.Context, tracer opentracing.Tracer, bookService book.Service, logger *log.Logger, fluentLogger *fluent.Fluent) *Server {
	server := &Server{
		ctx:          ctx,
		tracer:       tracer,
		bookService:  bookService,
		Logger:       logger,
		FluentLogger: fluentLogger,
	}

	bookHandler := bookHandler{ctx, tracer, bookService, logger, fluentLogger}

	router := mux.NewRouter()

	bookHandler.registerRouter(router)

	server.Router = router

	return server
}

// Run runs the HTTP server with the specified port and router.
func (srv *Server) Run(port string) {
	port = ":" + port

	srv.Logger.Log.WithFields(log.Fields{
		"service": "bookService",
		"port":    port,
	}).Info("bookService is running")
	srv.FluentLogger.Post("server", "bookService running on :8082")

	err := http.ListenAndServe(port, srv.Router)
	if err != nil {
		srv.Logger.Log.Panic(err)
	}
}
