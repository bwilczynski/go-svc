package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	h "github.com/bwilczynski/go-svc/http"
)

func main() {
	srv := &http.Server{
		Addr:    ":8090",
		Handler: h.NewService(),
	}
	go func() {
		srv.ListenAndServe()
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)
	s := <-done
	fmt.Printf("Received signal %v, performing graceful shutdown", s)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
}
