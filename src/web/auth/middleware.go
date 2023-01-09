package auth

import (
	"context"
	"github.com/mikejeuga/currency_converter/config"
	"net/http"
)

const TheApiKey = "X-Api-Key"

func NewMiddleware(config config.Config) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			headerVal := r.Header.Get(TheApiKey)
			if headerVal != config.ApiKey {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			r = r.WithContext(WithApiKey(r.Context(), APIKEY(headerVal)))
			next.ServeHTTP(w, r)
		})
	}
}

type APIKEY string

type contextKey struct{}

func LookUpApiKey(ctx context.Context) (APIKEY, bool) {
	apikey, ok := ctx.Value(contextKey{}).(APIKEY)
	return apikey, ok
}

func WithApiKey(ctx context.Context, apikey APIKEY) context.Context {
	return context.WithValue(ctx, contextKey{}, apikey)
}
