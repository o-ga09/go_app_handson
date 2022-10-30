package handler

import (
	"net/http"

	"github.com/taiti09/go_app_handson/auth"
)

func Authmiddleware(j *auth.JWTer) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			req, err := j.FillContext(r)
			if err != nil {
				RespondJSON(r.Context(),w,ErrResponse{
					Message: "not Found auth info",
					Details: []string{err.Error()},
				},http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w,req)
		})
	}
}

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !auth.IsAdmin(r.Context()) {
			RespondJSON(r.Context(),w,ErrResponse{
				Message: "not admin",
			},http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w,r)
	})
}