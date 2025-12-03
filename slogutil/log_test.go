package slogutil_test

import (
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
