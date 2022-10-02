package admin

import "github.com/bwilczynski/go-svc/pkg/http/metrics"

func (svc service) routes() {
	svc.mux.Handle("/health", svc.healthHandler())
	svc.mux.Handle("/info", svc.infoHandler())
	svc.mux.Handle("/metrics", metrics.Handler())
}
