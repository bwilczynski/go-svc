package http

import (
	"net/http"

	httpe "github.com/bwilczynski/go-svc/pkg/http"

	"github.com/bwilczynski/go-svc/pkg/http/metrics"
)

func (svc service) routes() {
	observe := httpe.NewMiddlewareChain(
		metrics.InstrumentHandler(func(r *http.Request) string { return svc.mux.GetRoutePattern(r) }),
		httpe.LoggingHandler(svc.logger),
	).Handler

	svc.mux.Handle("/hello", observe(svc.helloHandler()))
	svc.mux.Handle("/", observe(http.NotFoundHandler()))
}
