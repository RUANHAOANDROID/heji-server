package cmd

import (
	"heji-server/config"
	"heji-server/internal/server"
	"heji-server/pkg"
	"time"

	"os"
)

var log = pkg.Log

const appName = "heji"
const appAbout = "                                           \n    _/    _/  _/_/_/_/        _/  _/_/_/   \n   _/    _/  _/              _/    _/      \n  _/_/_/_/  _/_/_/          _/    _/       \n _/    _/  _/        _/    _/    _/        \n_/    _/  _/_/_/_/    _/_/    _/_/_/       \n                                           "
const appEdition = "ce"
const appDescription = "合記服务节点\n\t去中心化|多人同时记账|账单统计可视化|账单区分权限|账单导入|账单导出"
const appCopyright = "(c) 2023-2024 hao88.cloud. All rights reserved."

var version = "development"

func Main(args []string) {
	defer func() {
		if r := recover(); r != nil {
			log.Error(r)
			os.Exit(1)
		}
	}()
	conf := config.Load()
	start := time.Now()
	log.Info(start)
	log.Info(appName)
	log.Info(appAbout)
	log.Info(appEdition)
	log.Info(appDescription)
	log.Info(appCopyright)
	server.Start(conf)
}
