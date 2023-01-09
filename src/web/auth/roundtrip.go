package auth

import "net/http"

type MyRoundTripper struct {
	Next http.RoundTripper
}

func (m MyRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	ctx := r.Context()
	value, found := LookUpApiKey(ctx)
	if found {
		r.Header.Set(TheApiKey, string(value))
	}
	return m.Next.RoundTrip(r)
}
