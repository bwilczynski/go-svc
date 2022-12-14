package http

import (
	"net/http"

	"github.com/bwilczynski/go-svc/pkg/httpe"

	"github.com/rs/zerolog"
)

type service struct {
	mux    *httpe.ServeMux
	logger zerolog.Logger
}

func NewService(logger zerolog.Logger) *service {
	svc := &service{mux: &httpe.ServeMux{}, logger: logger}
	svc.routes()
	return svc
}

func (svc service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	svc.mux.ServeHTTP(w, r)
}
