package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	h "github.com/bwilczynski/go-svc/http"
	"github.com/rs/zerolog"
)

var (
	console = flag.Bool("console", false, "Enable pretty logging on console")
	debug   = flag.Bool("debug", false, "Sets log level to debug")
	port    = flag.Int("port", 8000, "HTTP port to run server on")
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running the server: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	flag.Parse()

	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Logger()
	if *console {
		logger = logger.Output(zerolog.NewConsoleWriter())
	}
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		logger.Warn().Msg("Server running in DEBUG mode.")
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	logger.Info().Msg("Application is starting")

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: h.NewService(logger),
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

	return nil
}
