// Package slogutil provides structured logging functionality for ghtkn.
// It uses slog with tint handler for colored output to stderr.
package slogutil

import (
	"errors"
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
	"github.com/mattn/go-colorable"
)

type InputNew struct {
	Name    string
	Version string
	Level   slog.Level
	Out     *os.File
}

// New creates a new structured logger with the specified version and log level.
// The logger outputs to stderr with colored formatting using tint handler.
// It includes "program" and "version" attributes in all log entries.
func New(input *InputNew) *slog.Logger {
	out := input.Out
	if out == nil {
		out = os.Stderr
	}
	return slog.New(tint.NewHandler(colorable.NewColorable(out), &tint.Options{
		Level: input.Level,
	})).With("program", input.Name, "version", input.Version)
}

// ErrUnknownLogLevel is returned when an invalid log level string is provided to ParseLevel.
var ErrUnknownLogLevel = errors.New("unknown log level")

func SetLevel(levelVar *slog.LevelVar, level string) error {
	lvl, err := parseLevel(level)
	if err != nil {
		return err
	}
	levelVar.Set(lvl)
	return nil
}

func parseLevel(lvl string) (slog.Level, error) {
	switch lvl {
	case "debug":
		return slog.LevelDebug, nil
	case "info":
		return slog.LevelInfo, nil
	case "warn":
		return slog.LevelWarn, nil
	case "error":
		return slog.LevelError, nil
	default:
		return 0, ErrUnknownLogLevel
	}
}
