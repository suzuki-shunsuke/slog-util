package slogutil

import (
	"errors"
)

type ColorMode int

const (
	ColorAuto ColorMode = iota
	ColorOn
	ColorOff
)

func (l *Logger) SetColor(color string) error {
	if color == "" {
		return nil
	}
	mode, err := parseColorMode(color)
	if err != nil {
		return err
	}
	switch mode {
	case ColorOn:
		l.tintOpts.NoColor = false
	case ColorOff:
		l.tintOpts.NoColor = true
	case ColorAuto:
		l.tintOpts.NoColor = !l.autoColor
	}
	return nil
}

func (l *Logger) EnableColor() {
	l.tintOpts.NoColor = false
}

func (l *Logger) DisableColor() {
	l.tintOpts.NoColor = true
}

var ErrUnknownColorMode = errors.New("unknown color mode")

func parseColorMode(color string) (ColorMode, error) {
	switch color {
	case "auto":
		return ColorAuto, nil
	case "on", "always":
		return ColorOn, nil
	case "off", "never":
		return ColorOff, nil
	default:
		return 0, ErrUnknownColorMode
	}
}
