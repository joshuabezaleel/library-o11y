package server

import (
	"log"
	"net/http"

	"github.com/joshuabezaleel/library-o11y/book"

	"github.com/gorilla/mux"
)

// Server holds dependencies for the Book HTTP server.
type Server struct {
	bookService book.Service

	Router *mux.Router
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

	return server
}

// Run runs the HTTP server with the specified port and router.
func (srv *Server) Run(port string) {
	port = ":" + port

	log.Println("bookService is running on port " + port)
	err := http.ListenAndServe(port, srv.Router)
	if err != nil {
		panic(err)
	}
}
