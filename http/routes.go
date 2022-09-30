package http

func (r service) routes() {
	r.mux.Handle("/hello", r.handleHello())
	r.mux.Handle("/healthz", r.handleHealthz())
}
