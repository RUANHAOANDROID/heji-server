package main

import (
	"github.com/urfave/cli"
	"heji-server/utils"
	"os"
	"sync"
)

var log = utils.Log

const appName = "heji"
const appAbout = "                                           \n    _/    _/  _/_/_/_/        _/  _/_/_/   \n   _/    _/  _/              _/    _/      \n  _/_/_/_/  _/_/_/          _/    _/       \n _/    _/  _/        _/    _/    _/        \n_/    _/  _/_/_/_/    _/_/    _/_/_/       \n                                           "
const appEdition = "ce"
const appDescription = "合記服务节点\n\t去中心化|多人同时记账|账单统计可视化|账单区分权限|账单导入|账单导出"
const appCopyright = "(c) 2023-2024 hao88.cloud. All rights reserved."

var version = "development"

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Error(r)
			os.Exit(1)
		}
	}()
	var wg sync.WaitGroup
	app := cli.NewApp()
	app.Name = appName
	app.Usage = appAbout
	app.Description = appDescription
	app.Version = version
	app.Copyright = appCopyright
	app.EnableBashCompletion = true
	app.Commands = Commands
	app.Action = func(ctx *cli.Context) {
		println("do action...")
		startAction(ctx)
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
	wg.Wait()
}
