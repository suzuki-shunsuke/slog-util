package slogutil

import (
	"errors"
	"log/slog"
)

// ErrUnknownLogLevel is returned when an invalid log level string is provided to ParseLevel.
var ErrUnknownLogLevel = errors.New("unknown log level")

func (l *Logger) SetLevel(level string) error {
	return setLevel(l.level, level)
}

func setLevel(levelVar *slog.LevelVar, level string) error {
	if level == "" {
		return nil
	}
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
