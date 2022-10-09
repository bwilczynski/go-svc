package http

import (
	"net/http"

	"github.com/bwilczynski/go-svc/pkg/httpe"

	"github.com/bwilczynski/go-svc/pkg/httpe/metrics"
)

func (svc service) routes() {
	transport := httpe.DumpRequestTransport(svc.logger)(http.DefaultTransport)

	svc.mux.Handle("/hello", svc.observe(svc.helloHandler()))
	svc.mux.Handle("/httpbin/", svc.observe(http.StripPrefix("/httpbin/", svc.httpbinHandler(transport))))
	svc.mux.Handle("/", svc.observe(http.NotFoundHandler()))
}

func (svc service) observe(h http.Handler) http.Handler {
	return httpe.NewMiddlewareChain(
		metrics.InstrumentHandler(func(r *http.Request) string { return svc.mux.GetRoutePattern(r) }),
		httpe.LoggingHandler(svc.logger),
		httpe.DumpRequestHandler(svc.logger),
	).Handler(h)
}
