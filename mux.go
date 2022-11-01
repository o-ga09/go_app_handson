package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/taiti09/go_app_handson/auth"
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
	clocker := clock.RealClocker{}
	r := store.Repository{Clocker: clocker}
	rchi, err := store.NewKVS(ctx,cfg)
	if err != nil {
		return nil,cleanup,err
	}
	jwter, err := auth.NewJWTer(rchi,clocker)
	if err != nil {
		return nil,cleanup,err
	}
	l := &handler.Login{
		Service: &service.Login{
			DB: db,
			Repo: &r,
			TokenGenerator: jwter,
		},
		Validator: v,
	}
	r = store.Repository{Clocker: clock.RealClocker{}}
	at := &handler.AddTask{
		Service: &service.AddTask{DB: db, Repo: &r},
		Validator: v,
	}
	lt := &handler.ListTask{
		Service: &service.ListTask{DB: db, Repo: &r},
		
	}

	mux.Post("/login",l.ServeHTTP)

	mux.Route("/tasks",func(r chi.Router){
		r.Use(handler.Authmiddleware(jwter))
        r.Get("/",lt.ServeHTTP)
        r.Post("/",at.ServeHTTP)
	})
	mux.Route("/admin",func(r chi.Router) {
		r.Use(handler.Authmiddleware(jwter),handler.AdminMiddleware)
		r.Get("/",func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Context-type","application/json; charset=utf-8")
			_, _ = w.Write([]byte(`{"meesage": "admin only"}`))
		})
	})
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