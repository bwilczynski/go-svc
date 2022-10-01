package http

import (
	"net/http"

	"github.com/rs/zerolog"
)

type service struct {
	mux    *http.ServeMux
	logger zerolog.Logger
}

func NewService(logger zerolog.Logger) *service {
	svc := &service{mux: http.NewServeMux(), logger: logger}
	svc.routes()
	return svc
}

func (r service) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}
