package slogutil

import (
	"io"
	"log/slog"
	"os"
)

type JSONLogger struct {
	logger *slog.Logger
	hOpts  *slog.HandlerOptions
	level  *slog.LevelVar
}

type InputNewJSON struct {
	Name           string
	Version        string
	Out            io.Writer
	HandlerOptions *slog.HandlerOptions
}

// NewJSON creates a new structured logger with the specified version and log level.
// The logger outputs to stderr with colored formatting using tint handler.
// It includes "program" and "version" attributes in all log entries.
func NewJSON(input *InputNewJSON) *JSONLogger {
	out := input.Out
	if out == nil {
		out = os.Stderr
	}
	levelVar := &slog.LevelVar{}
	if input.HandlerOptions == nil {
		input.HandlerOptions = &slog.HandlerOptions{}
	}
	input.HandlerOptions.Level = levelVar
	return &JSONLogger{
		logger: slog.New(slog.NewJSONHandler(out, input.HandlerOptions)).With("program", input.Name, "version", input.Version),
		hOpts:  input.HandlerOptions,
		level:  levelVar,
	}
}

func (l *JSONLogger) Logger() *slog.Logger {
	return l.logger
}

func (l *JSONLogger) SetLevel(level string) error {
	return setLevel(l.level, level)
}
