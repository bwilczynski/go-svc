package metrics

import (
	"net/http"

	"github.com/bwilczynski/go-svc/pkg/httpe"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var inFlightGauge = prometheus.NewGauge(prometheus.GaugeOpts{
	Name:      "requests_in_flight",
	Subsystem: "http_server",
	Help:      "A gauge of requests currently being served by the wrapped handler.",
})

var counter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name:      "requests_total",
		Subsystem: "http_server",
		Help:      "A counter for requests to the wrapped handler.",
	},
	[]string{"code", "method"},
)

var duration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:      "requests_seconds",
		Subsystem: "http_server",
		Help:      "A histogram of latencies for HTTP requests.",
	},
	[]string{"url", "code", "method"},
)

var responseSize = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:      "response_size_bytes",
		Subsystem: "http_server",
		Help:      "A histogram of response sizes for requests.",
		Buckets:   []float64{200, 500, 900, 1500},
	},
	[]string{},
)

func init() {
	prometheus.MustRegister(inFlightGauge, counter, duration, responseSize)
}

func InstrumentHandler(urlFunc func(r *http.Request) string) httpe.MiddlewareFunc[http.Handler] {
	chain := httpe.NewMiddlewareChain(
		func(next http.Handler) http.Handler {
			return promhttp.InstrumentHandlerResponseSize(responseSize, next)
		},
		func(next http.Handler) http.Handler {
			return promhttp.InstrumentHandlerCounter(counter, next)
		},
		func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				promhttp.InstrumentHandlerDuration(
					duration.MustCurryWith(prometheus.Labels{"url": urlFunc(r)}), next)(w, r)
			})
		},
		func(next http.Handler) http.Handler {
			return promhttp.InstrumentHandlerInFlight(inFlightGauge, next)
		},
	)
	return chain.Handler
}

func Handler() http.Handler {
	return promhttp.Handler()
}
