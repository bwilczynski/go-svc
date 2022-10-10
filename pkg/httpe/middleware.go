package httpe

import (
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
)

type MiddlewareFunc[T any] func(T) T

type middlewareChain[T any] struct {
	items []MiddlewareFunc[T]
}

func NewMiddlewareChain[T any](items ...MiddlewareFunc[T]) middlewareChain[T] {
	return middlewareChain[T]{items: items}
}

func (chain middlewareChain[T]) Handler(next T) T {
	for i := len(chain.items) - 1; i >= 0; i-- {
		next = chain.items[i](next)
	}
	return next
}

func logAccess(r *http.Request, msg string, status, size int, duration time.Duration) {
	hlog.FromRequest(r).Info().
		Str("method", r.Method).
		Stringer("url", r.URL).
		Str("content_type", r.Header.Get("Content-Type")).
		Int("status", status).
		Int("size", size).
		Dur("duration", duration).
		Msg(msg)
}

func LoggingHandler(logger zerolog.Logger) MiddlewareFunc[http.Handler] {
	chain := NewMiddlewareChain(
		hlog.NewHandler(logger),
		hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
			logAccess(r, "Access", status, size, duration)
		}),
		hlog.RemoteAddrHandler("ip"),
		hlog.UserAgentHandler("user_agent"),
		hlog.RequestIDHandler("req_id", "Request-Id"),
	)
	return chain.Handler
}

func DumpRequestHandler() MiddlewareFunc[http.Handler] {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			logger := hlog.FromRequest(req)
			if logger.Debug().Enabled() {
				if r, err := httputil.DumpRequest(req, true); err == nil {
					logger.Debug().
						Str("req", string(r)).
						Msg("Access")
				}
			}
			next.ServeHTTP(w, req)
		})
	}
}

type RoundTripperFunc func(*http.Request) (*http.Response, error)

func (f RoundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

func LoggingTransport() MiddlewareFunc[http.RoundTripper] {
	return func(next http.RoundTripper) http.RoundTripper {
		return RoundTripperFunc(func(r *http.Request) (*http.Response, error) {
			start := time.Now()
			res, err := next.RoundTrip(r)
			logAccess(r, "HTTP request", res.StatusCode, int(res.ContentLength), time.Since(start))
			return res, err
		})
	}
}

func DumpRequestTransport() MiddlewareFunc[http.RoundTripper] {
	return func(next http.RoundTripper) http.RoundTripper {
		return RoundTripperFunc(func(req *http.Request) (*http.Response, error) {
			logger := hlog.FromRequest(req)
			if logger.Debug().Enabled() {
				if r, err := httputil.DumpRequest(req, true); err == nil {
					logger.Debug().
						Str("req", string(r)).
						Msg("HTTP request")
				}
			}
			return next.RoundTrip(req)
		})
	}
}
