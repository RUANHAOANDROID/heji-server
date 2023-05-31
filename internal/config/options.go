package config

import "github.com/urfave/cli"

type Options struct {
	AppName        string   `yaml:"AppName" json:"AppName" flag:"AppName"`
	Prod           bool     `yaml:"Prod" json:"Prod" flag:"prod"`
	Debug          bool     `yaml:"Debug" json:"Debug" flag:"debug"`
	Version        string   `json:"-"`
	TrustedProxies []string `yaml:"TrustedProxies" json:"-" flag:"trusted-proxy"`
}

func NewOptions(ctx *cli.Context) *Options {
	c := &Options{}
	if ctx == nil {
		return c
	}
	c.AppName = "heji-server"
	c.Debug = true
	c.Version = ctx.App.Version
	return c
}
