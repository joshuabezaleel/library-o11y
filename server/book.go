package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/joshuabezaleel/library-o11y/book"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type bookHandler struct {
	bookService book.Service
}

func (handler *bookHandler) registerRouter(router *mux.Router) {
	router.HandleFunc("/books", handler.getAllBooks).Methods("GET")
	router.HandleFunc("/books/{ID}", handler.getBook).Methods("GET")
	router.HandleFunc("/metrics", handler.metrics).Methods("GET")
}

func (handler *bookHandler) metrics(w http.ResponseWriter, r *http.Request) {
	promhttp.Handler().ServeHTTP(w, r)
}

func (handler *bookHandler) getAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := handler.bookService.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// getAllBooksCounter := promauto.NewCounter(prometheus.CounterOpts{
	// 	Name: "getAllBooks_total",
	// 	Help: "Total number of processed getAllBooks hit",
	// })
	// getAllBooksCounter.Inc()

	respondWithJSON(w, http.StatusOK, books)
}

func (handler *bookHandler) getBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookIDString, ok := vars["ID"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, "Invalid URL Path")
		return
	}

	bookID, _ := strconv.Atoi(bookIDString)

	retrievedBook, err := handler.bookService.Get(bookID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, retrievedBook)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"Error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
