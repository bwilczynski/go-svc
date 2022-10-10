package httpe

import (
	"log"
	"strings"

	"github.com/rs/zerolog"
)

type zerologWriter struct {
	logger zerolog.Logger
}

func NewLogger(logger zerolog.Logger) *log.Logger {
	return log.New(&zerologWriter{logger: logger}, "", 0)
}

func (w zerologWriter) Write(p []byte) (n int, err error) {
	w.logger.Error().Msg(strings.TrimSpace(string(p)))
	return len(p), nil
}
