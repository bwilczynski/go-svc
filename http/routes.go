package http

func (svc service) routes() {
	svc.mux.Handle("/hello", withLogger(svc.logger)(svc.handleHello()))
}
