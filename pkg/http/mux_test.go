package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeMux(t *testing.T) {
	mux := ServeMux{}
	pattern := "/expected"
	mux.Handle(pattern, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if p := mux.GetRoutePattern(r); p != pattern {
			t.Fatalf("route pattern should be: %q, got %q", pattern, p)
		}
	}))

	w := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, pattern, nil)
	if err != nil {
		t.Fatalf("cannot create request: %v", err)
	}
	mux.ServeHTTP(w, r)
}
