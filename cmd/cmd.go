package cmd

import (
	"context"
	"github.com/urfave/cli"
	"heji-server/internal/server"
)

var Commands = []cli.Command{
	StartCommand,
}
var StartCommand = cli.Command{
	Name:    "start",
	Aliases: []string{"up"},
	Usage:   "Start server",
	Flags:   startFlags,
	Action:  startAction,
}

func startAction(ctx *cli.Context) error {
	conf, err := InitConfig(ctx)
	if err != nil {
		return err
	}
	cctx, _ := context.WithCancel(context.Background())
	server.Start(cctx, conf)
	return nil
}

// startFlags specifies the start command parameters.
var startFlags = []cli.Flag{
	cli.BoolFlag{
		Name:   "detach-server, d",
		Usage:  "detach from the console (daemon mode)",
		EnvVar: "PHOTOPRISM_DETACH_SERVER",
	},
	cli.BoolFlag{
		Name:  "config, c",
		Usage: "show config",
	},
}
