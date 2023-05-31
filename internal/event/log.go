package event

import (
	"github.com/leandro-lugaresi/hub"
	"github.com/sirupsen/logrus"
)

// TextFormatter for log messages.
var TextFormatter = &logrus.TextFormatter{
	DisableColors: false,
	FullTimestamp: true,
}

// Log is the global default logger.
var Log Logger
var LogBuffer Buffer

// Hook represents a log event hook.
type Hook struct {
	hub *hub.Hub
}

// NewHook creates a new log event hook.
func NewHook(hub *hub.Hub) *Hook {
	return &Hook{hub: hub}
}

// Levels returns a slice containing all supported log levels.
func (h *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}
