package main

import (
	"fmt"
	"github.com/urfave/cli"
	"heji-server/utils"
	"os"
)

var log = utils.Log

const appName = "HeJi"
const appAbout = "HeJi®"
const appEdition = "ce"
const appDescription = "heji-server"
const appCopyright = "(c) 2023-2024 HeJi UG. All rights reserved."

var version = "development"

func main() {
	log.Println("Start heji server")
	pinter()
	defer func() {
		if r := recover(); r != nil {
			log.Error(r)
			os.Exit(1)
		}
	}()

	app := cli.NewApp()
	app.Name = appName
	app.Usage = appAbout
	app.Description = appDescription
	app.Version = version
	app.Copyright = appCopyright
	app.EnableBashCompletion = true
	app.Commands = Commands
	app.Action = func(*cli.Context) error {
		fmt.Println("cli do action!")
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
func pinter() {
	log.Println("this is pinter func")
}
