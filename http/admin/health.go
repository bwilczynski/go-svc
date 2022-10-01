package admin

import (
	"net/http"

	httpe "github.com/bwilczynski/go-svc/pkg/http"
)

func (svc service) healthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpe.JSON(w, map[string]string{"status": "ok"})
	}
}
