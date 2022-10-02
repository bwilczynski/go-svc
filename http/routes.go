package http

import (
	"net/http"

	httpe "github.com/bwilczynski/go-svc/pkg/http"

	"github.com/bwilczynski/go-svc/pkg/http/metrics"
)

func (svc service) routes() {
	svc.handle("/hello", svc.helloHandler())
	svc.handle("/", http.NotFoundHandler())
}

func (svc service) handle(pattern string, handler http.Handler) {
	observe := httpe.NewMiddlewareChain(
		metrics.InstrumentHandler(func(r *http.Request) string { return svc.mux.GetRoutePattern(r) }),
		httpe.LoggingHandler(svc.logger),
	).Handler
	svc.mux.Handle(pattern, observe(handler))
}
