package slogtest

import (
	"log/slog"
	"testing"
)

func TestNewNopHandler(t *testing.T) {
	l := slog.New(NewNopHandler())
	l.Info("hello")
}

func TestNopHandler_Enabled(t *testing.T) {
	levels := []slog.Level{
		slog.LevelDebug,
		slog.LevelInfo,
		slog.LevelWarn,
		slog.LevelError,
	}

	l := slog.New(NewNopHandler())
	for _, level := range levels {
		if l.Enabled(nil, level) {
			t.Errorf("expected false for level %s", level)
		}
	}
}

func TestNopHandler_WithAttrs(t *testing.T) {
	l := slog.New(NewNopHandler())
	l = l.With(slog.String("foo", "bar"))
	l.Info("hello")
}

func TestNopHandler_WithGroup(t *testing.T) {
	l := slog.New(NewNopHandler())
	l = l.WithGroup("foo")
	l.Info("hello")
}
