package http

import (
	"net/http"

	"github.com/bwilczynski/go-svc/pkg/metrics"
)

func (svc service) routes() {
	observe := func(next http.Handler) http.Handler {
		m := metrics.InstrumentHandler(func(r *http.Request) string { return svc.mux.GetRoutePattern(r) })
		l := loggingHandler(svc.logger)(next)
		return m(l)
	}
	svc.mux.Handle("/hello", observe(svc.helloHandler()))
	svc.mux.Handle("/", observe(http.NotFoundHandler()))
}
