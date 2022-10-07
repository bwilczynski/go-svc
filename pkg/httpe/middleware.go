package httpe

import (
	"net/http"
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

func LoggingHandler(logger zerolog.Logger) MiddlewareFunc[http.Handler] {
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
