package http

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func (svc service) httpbinHandler(transport http.RoundTripper) http.Handler {
	url, _ := url.Parse("https://httpbin.org")
	return svc.proxyHandler(url, transport)
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
