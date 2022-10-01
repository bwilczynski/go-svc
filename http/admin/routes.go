package admin

func (svc service) routes() {
	svc.mux.Handle("/healthz", svc.handleHealthz())
}
