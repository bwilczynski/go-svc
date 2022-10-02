package admin

import (
	"net/http"

	"github.com/bwilczynski/go-svc/pkg/httpe"
	"github.com/bwilczynski/go-svc/pkg/version"
)

func (svc service) infoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		info := version.Get()
		httpe.JSON(w, info)
	}
}
