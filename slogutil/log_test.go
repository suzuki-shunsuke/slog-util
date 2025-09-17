package slogutil_test

import (
	"errors"
	"log/slog"
	"testing"

	"github.com/suzuki-shunsuke/slog-util/slogutil"
)

func TestNew(t *testing.T) {
	t.Parallel()
	logger := slogutil.New(&slogutil.InputNew{
		Name:    "test",
		Version: "v0.0.1",
	})
	if logger == nil {
		t.Fatal("New() returned nil logger")
	}
}

func TestParseLevel(t *testing.T) { //nolint:funlen
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
			wantErr: slogutil.ErrUnknownLogLevel,
		},
		{
			name:    "empty string",
			input:   "",
			want:    0,
			wantErr: slogutil.ErrUnknownLogLevel,
		},
		{
			name:    "uppercase level",
			input:   "DEBUG",
			want:    0,
			wantErr: slogutil.ErrUnknownLogLevel,
		},
		{
			name:    "mixed case level",
			input:   "Info",
			want:    0,
			wantErr: slogutil.ErrUnknownLogLevel,
		},
		{
			name:    "level with spaces",
			input:   " info ",
			want:    0,
			wantErr: slogutil.ErrUnknownLogLevel,
		},
		{
			name:    "numeric level",
			input:   "0",
			want:    0,
			wantErr: slogutil.ErrUnknownLogLevel,
		},
		{
			name:    "verbose level not supported",
			input:   "verbose",
			want:    0,
			wantErr: slogutil.ErrUnknownLogLevel,
		},
		{
			name:    "trace level not supported",
			input:   "trace",
			want:    0,
			wantErr: slogutil.ErrUnknownLogLevel,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := slogutil.ParseLevel(tt.input)

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

func TestErrUnknownLogLevel(t *testing.T) {
	t.Parallel()
	// Test that the error has the expected message
	if got := slogutil.ErrUnknownLogLevel.Error(); got != "unknown log level" {
		t.Errorf("ErrUnknownLogLevel.Error() = %q, want %q", got, "unknown log level")
	}
}
