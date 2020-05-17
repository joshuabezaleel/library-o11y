package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/joshuabezaleel/library-o11y/book"
	"github.com/joshuabezaleel/library-o11y/log"

	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/gorilla/mux"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	getAllBooksCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "getAllBooks_total",
		Help: "Total number of processed getAllBooks hit",
	})
)

type bookHandler struct {
	ctx    context.Context
	tracer opentracing.Tracer

	bookService book.Service

	logger *log.Logger

	fluentLogger *fluent.Fluent
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
	spanCtx, _ := handler.tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	span := handler.tracer.StartSpan("getAllBooks", ext.RPCServerOption(spanCtx))
	defer span.Finish()

	ctx := r.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	books, err := handler.bookService.GetAll(ctx)
	if err != nil {
		handler.logger.Log.Debug("Error GET /books")
		handler.fluentLogger.Post("bookHandler", "Error GET /books")
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	handler.logger.Log.Info("GET /books")
	handler.fluentLogger.Post("bookHandler", "GET /books")
	span.LogKV("bookHandler", "getAllBooks")

	getAllBooksCounter.Inc()

	respondWithJSON(w, http.StatusOK, books)
}

func (handler *bookHandler) getBook(w http.ResponseWriter, r *http.Request) {
	spanCtx, _ := handler.tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	span := handler.tracer.StartSpan("getBook", ext.RPCServerOption(spanCtx))
	defer span.Finish()

	ctx := r.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	vars := mux.Vars(r)
	bookIDString, ok := vars["ID"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, "Invalid URL Path")
		return
	}

	bookID, _ := strconv.Atoi(bookIDString)

	logMessage := fmt.Sprintf("GET /books/%v", bookID)
	handler.logger.Log.Infof(logMessage)
	handler.fluentLogger.Post("bookHandler", logMessage)
	span.LogKV("bookHandler", "getBook")

	retrievedBook, err := handler.bookService.Get(ctx, bookID)
	if err != nil {
		errorLogMessage := fmt.Sprintf("Error GET /books/%v", bookID)

		handler.logger.Log.Debugf(errorLogMessage)
		handler.fluentLogger.Post("bookHandler", errorLogMessage)

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
