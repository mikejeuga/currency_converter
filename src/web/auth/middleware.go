package auth

import (
	"github.com/mikejeuga/currency_converter/config"
	"net/http"
)

func NewMiddleware(config config.Config) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			headerVal := r.Header.Get("X-Api-Key")
			if headerVal != config.ApiKey {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
