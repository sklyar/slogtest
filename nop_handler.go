package slogtest

import (
	"context"
	"log/slog"
)

// NopHandler is a Handler that does nothing.
type NopHandler struct{}

var _ slog.Handler = (*NopHandler)(nil)

// NewNopHandler creates a NopHandler.
func NewNopHandler() *NopHandler {
	return &NopHandler{}
}

// Enabled reports whether the handler handles records at the given level.
// Always returns false.
func (NopHandler) Enabled(context.Context, slog.Level) bool {
	return false
}

// Handle does nothing.
func (NopHandler) Handle(context.Context, slog.Record) error {
	return nil
}

// WithAttrs returns a new NopHandler.
func (NopHandler) WithAttrs([]slog.Attr) slog.Handler {
	return &NopHandler{}
}

// WithGroup returns a new NopHandler.
func (NopHandler) WithGroup(string) slog.Handler {
	return &NopHandler{}
}
