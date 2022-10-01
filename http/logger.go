package http

import (
	"net/http"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
)

func withLogger(logger zerolog.Logger) func(next http.Handler) http.Handler {
	h := hlog.NewHandler(logger)
	a := hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Stringer("url", r.URL).
			Str("content_type", r.Header.Get("Content-Type")).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Msg("")
	})
	ra := hlog.RemoteAddrHandler("ip")
	ua := hlog.UserAgentHandler("user_agent")
	rid := hlog.RequestIDHandler("req_id", "Request-Id")
	return func(next http.Handler) http.Handler {
		return h(a(ra(ua(rid(next)))))
	}
}
