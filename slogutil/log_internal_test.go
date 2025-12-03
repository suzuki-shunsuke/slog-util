package slogutil

import (
	"errors"
	"log/slog"
	"testing"
)

func Test_parseLevel(t *testing.T) { //nolint:funlen
	t.Parallel()
	tests := []struct {
		name    string
		input   string
		want    slog.Level
		wantErr error
	}{
		{
			name:    "parse debug level",
			input:   "debug",
			want:    slog.LevelDebug,
			wantErr: nil,
		},
		{
			name:    "parse info level",
			input:   "info",
			want:    slog.LevelInfo,
			wantErr: nil,
		},
		{
			name:    "parse warn level",
			input:   "warn",
			want:    slog.LevelWarn,
			wantErr: nil,
		},
		{
			name:    "parse error level",
			input:   "error",
			want:    slog.LevelError,
			wantErr: nil,
		},
		{
			name:    "unknown level",
			input:   "unknown",
			want:    0,
			wantErr: ErrUnknownLogLevel,
		},
		{
			name:    "empty string",
			input:   "",
			want:    0,
			wantErr: ErrUnknownLogLevel,
		},
		{
			name:    "uppercase level",
			input:   "DEBUG",
			want:    0,
			wantErr: ErrUnknownLogLevel,
		},
		{
			name:    "mixed case level",
			input:   "Info",
			want:    0,
			wantErr: ErrUnknownLogLevel,
		},
		{
			name:    "level with spaces",
			input:   " info ",
			want:    0,
			wantErr: ErrUnknownLogLevel,
		},
		{
			name:    "numeric level",
			input:   "0",
			want:    0,
			wantErr: ErrUnknownLogLevel,
		},
		{
			name:    "verbose level not supported",
			input:   "verbose",
			want:    0,
			wantErr: ErrUnknownLogLevel,
		},
		{
			name:    "trace level not supported",
			input:   "trace",
			want:    0,
			wantErr: ErrUnknownLogLevel,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := parseLevel(tt.input)

			if tt.wantErr != nil {
				if err == nil {
					t.Fatalf("ParseLevel() error = nil, wantErr %v", tt.wantErr)
				}
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("ParseLevel() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Fatalf("ParseLevel() unexpected error = %v", err)
			}

			if tt.want != got {
				t.Errorf("ParseLevel() mismatch (-want +got):\n- %v\n+ %v", tt.want, got)
			}
		})
	}
}
