package middlewares

import (
	"errors"
	"net/http"

	"github.com/sunaryaagung95/go-message-api/auth"
	"github.com/sunaryaagung95/go-message-api/responses"
)

// SetJSON response
func SetJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

// SetAuth func
func SetAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.CheckToken(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Sorry you have to login first"))
			return
		}
		next(w, r)
	}
}
