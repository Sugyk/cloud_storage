package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func Register(r *mux.Router, lg *logrus.Logger, db *sqlx.DB) {
	handler := newHandler(lg, db)
	// adding logger middleware
	r.Use(handler.MiddlewareLogger())
	r.HandleFunc("/healthz", handler.Health())
	r.HandleFunc("/todo", handler.Auth()).Methods(http.MethodPost)
	r.HandleFunc("/todo/{id}", handler.GetFile()).Methods(http.MethodGet)
	// r.HandleFunc("/todo/{id}", handler.Update()).Methods(http.MethodPut)
	r.HandleFunc("/todo/{id}", handler.Delete()).Methods(http.MethodDelete)
}
