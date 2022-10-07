package http

import (
	"net/http"

	"github.com/bwilczynski/go-svc/pkg/httpe"

	"github.com/bwilczynski/go-svc/pkg/httpe/metrics"
)

func (svc service) routes() {
	transport := LoggingRoundTripper(svc.logger)(http.DefaultTransport)

	svc.handle("/hello", svc.helloHandler())
	svc.handle("/httpbin/", http.StripPrefix("/httpbin/", svc.httpbinHandler(transport)))
	svc.handle("/", http.NotFoundHandler())
}

func (svc service) handle(pattern string, handler http.Handler) {
	observe := httpe.NewMiddlewareChain(
		metrics.InstrumentHandler(func(r *http.Request) string { return svc.mux.GetRoutePattern(r) }),
		httpe.LoggingHandler(svc.logger),
		httpe.DumpRequestHandler(svc.logger),
	).Handler
	svc.mux.Handle(pattern, observe(handler))
}
