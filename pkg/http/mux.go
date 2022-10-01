package http

import (
	"context"
	"net/http"
)

type ServeMux struct {
	http.ServeMux
}

type patternKey struct{}

func (mux *ServeMux) Handle(pattern string, handler http.Handler) {
	contextHandler := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), patternKey{}, pattern))
			next.ServeHTTP(w, r)
		})
	}
	mux.ServeMux.Handle(pattern, contextHandler(handler))
}

func (mux *ServeMux) GetRoutePattern(req *http.Request) string {
	return req.Context().Value(patternKey{}).(string)
}
