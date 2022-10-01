package http

import (
	"fmt"
	"net/http"
)

func (svc service) helloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello\n")
	}
}
