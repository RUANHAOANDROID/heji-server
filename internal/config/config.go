package config

import (
	"gorm.io/gorm"
	"sync"
)

// Config app config
type Config struct {
	once    sync.Once
	model   string
	db      *gorm.DB
	options Options
}

// Prod is prod
func (c *Config) Prod() bool {
	return c.options.Prod
}

// TrustedProxies returns proxy server ranges from which reverse proxy headers can be trusted.
func (c *Config) TrustedProxies() []string {
	return c.options.TrustedProxies
}
