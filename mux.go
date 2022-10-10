package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/taiti09/go_app_handson/handler"
	"github.com/taiti09/go_app_handson/store"
)

func NewMux() http.Handler {
	mux := chi.NewRouter()
	mux.HandleFunc("/health", HealthCheckHandler)
	v := validator.New()
	at := &handler.AddTask{Store: store.Tasks, Validator: v}
	mux.Post("/tasks",at.ServerHTTP)
	lt := &handler.ListTask{Store: store.Tasks}
	mux.Get("/tasks",lt.ServeHTTP)
	return mux
}

func HealthCheckHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type","application/json; charset=utf-8")
	_, _ = w.Write([]byte(`{"status": "ok"}`))
}