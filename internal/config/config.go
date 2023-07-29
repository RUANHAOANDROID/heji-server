package config

import (
	"github.com/urfave/cli"
	"gorm.io/gorm"
	"heji-server/internal/utils"
	"sync"
)

var log = utils.Log

// Config app config
type Config struct {
	once    sync.Once
	cliCtx  *cli.Context
	model   string
	db      *gorm.DB
	options *Options
}

// Prod is prod
func (c *Config) Prod() bool {
	return c.options.Prod
}

// TrustedProxies returns proxy server ranges from which reverse proxy headers can be trusted.
func (c *Config) TrustedProxies() []string {
	return c.options.TrustedProxies
}

func NewConfig(ctx *cli.Context) *Config {
	// Initialize options from config file and CLI context.
	c := &Config{
		cliCtx:  ctx,
		options: NewOptions(ctx),
	}
	return c
}
func (c *Config) Init() error {
	log.Debugln("config init:", c.model, c.options)
	return nil
}
