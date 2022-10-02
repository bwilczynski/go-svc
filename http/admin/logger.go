package admin

import (
	"fmt"
	"net/http"

	httpe "github.com/bwilczynski/go-svc/pkg/http"

	"github.com/rs/zerolog"
)

func (svc service) loggerHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			svc.updateLogger(w, r)
		case http.MethodGet:
			svc.getLogger(w, r)
		default:
			http.Error(w, "Not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func (svc service) updateLogger(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
	l, err := zerolog.ParseLevel(r.PostForm.Get("level"))
	if err != nil {
		svc.logger.Err(fmt.Errorf("parseLevel: %w", err)).Msg("")
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	svc.logger.Info().Msgf("Setting error level to %s", l)
	zerolog.SetGlobalLevel(l)
	w.WriteHeader(http.StatusNoContent)
}

func (svc service) getLogger(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	httpe.JSON(w, map[string]string{"level": zerolog.GlobalLevel().String()})
}
