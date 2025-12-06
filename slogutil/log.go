// Package slogutil provides structured logging functionality for ghtkn.
// It uses slog with tint handler for colored output to stderr.
package slogutil

import (
	"io"
	"log/slog"
	"os"
	"runtime"

	"github.com/lmittmann/tint"
	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
)

type Logger struct {
	*slog.Logger

	level     *slog.LevelVar
	tintOpts  *tint.Options
	autoColor bool
}

type InputNew struct {
	Name        string
	Version     string
	Out         *os.File
	TintOptions *tint.Options
}

// New creates a new structured logger with the specified version and log level.
// The logger outputs to stderr with colored formatting using tint handler.
// It includes "program" and "version" attributes in all log entries.
func New(input *InputNew) *Logger {
	out := input.Out
	if out == nil {
		out = os.Stderr
	}
	levelVar := &slog.LevelVar{}
	if input.TintOptions == nil {
		input.TintOptions = &tint.Options{}
	}
	input.TintOptions.Level = levelVar
	autoColor := isatty.IsTerminal(out.Fd())
	var w io.Writer = out
	if runtime.GOOS == "windows" {
		w = colorable.NewColorable(out)
	}
	return &Logger{
		Logger:    slog.New(tint.NewHandler(w, input.TintOptions)).With("program", input.Name, "version", input.Version),
		level:     levelVar,
		tintOpts:  input.TintOptions,
		autoColor: autoColor,
	}
}
