package main

import (
	"github.com/urfave/cli"
	"heji-server/internal/config"
	"heji-server/internal/get"
)

var InitConfig = func(ctx *cli.Context) (*config.Config, error) {
	c := config.NewConfig(ctx)
	get.SetConfig(c)
	return c, c.Init()
}
