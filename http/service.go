package http

import (
	"net/http"
)

type service struct {
	mux *http.ServeMux
}

func NewService() *service {
	svc := &service{mux: http.NewServeMux()}
	svc.routes()
	return svc
}

func (r service) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}
