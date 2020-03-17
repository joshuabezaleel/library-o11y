package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/joshuabezaleel/library-o11y/book"

	"github.com/gorilla/mux"
)

type bookHandler struct {
	bookService book.Service
}

func (handler *bookHandler) registerRouter(router *mux.Router) {
	router.HandleFunc("/books", handler.getAllBooks).Methods("GET")
}

func (handler *bookHandler) getAllBooks(w http.ResponseWriter, r *http.Request) {
	var topicIDs []int

	queryValues := r.URL.Query()

	topics := queryValues["topics"]
	if topics != nil {
		// topicIDsString := strings.Split(topics, " ")
		for _, topicIDString := range topics {
			topicID, _ := strconv.Atoi(topicIDString)

			topicIDs = append(topicIDs, topicID)
		}
	} else {
		topicIDs = []int{1, 2, 3, 4}
	}

	var sort string
	sortParam, ok := queryValues["sort"]
	if !ok {
		sort = "Title"
	} else {
		sort = sortParam[0]
	}

	var order string
	orderParam, ok := queryValues["order"]
	if !ok {
		order = "asc"
	} else {
		order = orderParam[0]
	}

	books, err := handler.bookService.GetAll(topicIDs, sort, order)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, books)
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
