package server

import (
	"net/http"

	"github.com/joshuabezaleel/library-o11y/book"
	"github.com/joshuabezaleel/library-o11y/log"

	"github.com/gorilla/mux"
)

// Server holds dependencies for the Book HTTP server.
type Server struct {
	bookService book.Service

	Router *mux.Router

	Logger *log.Logger
}

// NewServer returns a new Book HTTP server
// with all of the necessary dependencies.
func NewServer(bookService book.Service) *Server {
	server := &Server{
		bookService: bookService,
	}

	bookHandler := bookHandler{bookService}

	router := mux.NewRouter()

	bookHandler.registerRouter(router)

	server.Router = router

	server.Logger = log.NewLogger()

	return server
}

// Run runs the HTTP server with the specified port and router.
func (srv *Server) Run(port string) {
	port = ":" + port

	// srv.Logger.Info("bookService is running on port " + port)
	srv.Logger.Log.WithFields(log.Fields{
		"service": "bookService",
		"port":    port,
	}).Info("bookService is running")
	err := http.ListenAndServe(port, srv.Router)
	if err != nil {
		panic(err)
	}
}
