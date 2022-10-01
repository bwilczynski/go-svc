package admin

import "github.com/bwilczynski/go-svc/pkg/metrics"

func (svc service) routes() {
	svc.mux.Handle("/healthz", svc.healthHandler())
	svc.mux.Handle("/metrics", metrics.Handler())
}
