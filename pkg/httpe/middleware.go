package httpe

import (
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
)

type MiddlewareFunc func(http.Handler) http.Handler

type middlewareChain struct {
	items []MiddlewareFunc
}

func NewMiddlewareChain(items ...MiddlewareFunc) middlewareChain {
	return middlewareChain{items: items}
}

func (chain middlewareChain) Handler(next http.Handler) http.Handler {
	for i := len(chain.items) - 1; i >= 0; i-- {
		next = chain.items[i](next)
	}
	return next
}

func LoggingHandler(logger zerolog.Logger) MiddlewareFunc {
	chain := NewMiddlewareChain(
		hlog.NewHandler(logger),
		hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
			hlog.FromRequest(r).Info().
				Str("method", r.Method).
				Stringer("url", r.URL).
				Str("content_type", r.Header.Get("Content-Type")).
				Int("status", status).
				Int("size", size).
				Dur("duration", duration).
				Msg("")
		}),
		hlog.RemoteAddrHandler("ip"),
		hlog.UserAgentHandler("user_agent"),
		hlog.RequestIDHandler("req_id", "Request-Id"),
	)
	return chain.Handler
}

func DumpRequestHandler(logger zerolog.Logger) MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if logger.Debug().Enabled() {
				if r, err := httputil.DumpRequest(r, true); err == nil {
					logger.Debug().Msg(string(r))
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}
