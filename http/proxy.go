package http

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/rs/zerolog"
)

func (svc service) httpbinHandler(transport http.RoundTripper) http.Handler {
	url, _ := url.Parse("https://httpbin.org")
	return svc.proxyHandler(url, transport)
}

type RoundTripperFunc func(*http.Request) (*http.Response, error)

func (f RoundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

func LoggingRoundTripper(logger zerolog.Logger) func(next http.RoundTripper) http.RoundTripper {
	return func(next http.RoundTripper) http.RoundTripper {
		return RoundTripperFunc(func(r *http.Request) (*http.Response, error) {
			if logger.Debug().Enabled() {
				if r, err := httputil.DumpRequest(r, true); err == nil {
					logger.Debug().Msg(string(r))
				}
			}
			return next.RoundTrip(r)
		})
	}
}

func (svc service) proxyHandler(target *url.URL, transport http.RoundTripper) http.Handler {
	rp := httputil.NewSingleHostReverseProxy(target)
	pass := rp.Director
	rp.Director = func(r *http.Request) {
		pass(r)
		r.Host = target.Host
	}
	rp.Transport = transport
	return rp
}
