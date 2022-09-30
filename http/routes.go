package http

func (r service) routes() {
	r.mux.Handle("/hello", r.handleHello())
}
