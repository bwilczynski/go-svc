package http

import "github.com/bwilczynski/go-svc/pkg/metrics"

func (svc service) routes() {
	svc.mux.Handle("/hello", metrics.InstrumentHandler("/hello")(loggingHandler(svc.logger)(svc.helloHandler())))
}
