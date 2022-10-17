package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/taiti09/go_app_handson/clock"
	"github.com/taiti09/go_app_handson/config"
	"github.com/taiti09/go_app_handson/handler"
	"github.com/taiti09/go_app_handson/service"
	"github.com/taiti09/go_app_handson/store"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	mux := chi.NewRouter()
	mux.HandleFunc("/health", HealthCheckHandler)
	v := validator.New()
	db, cleanup, err := store.New(ctx,cfg)
	if err != nil {
		return nil, cleanup, err
	}
	r := store.Repository{Clocker: clock.RealClocker{}}
	at := &handler.AddTask{
		Service: &service.AddTask{DB: db, Repo: &r},
		Validator: v,
	}
	mux.Post("/tasks",at.ServeHTTP)
	lt := &handler.ListTask{
		Service: &service.ListTask{DB: db, Repo: &r},
		
	}
	mux.Get("/tasks",lt.ServeHTTP)
	ru := &handler.RegisterUser{
		Service: &service.RegisterUser{DB: db, Repo: &r},
		Validator: v,
	}
	mux.Post("/register",ru.ServeHTTP)
	return mux, cleanup, nil
}

func HealthCheckHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type","application/json; charset=utf-8")
	_, _ = w.Write([]byte(`{"status": "ok"}`))
}