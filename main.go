package main

import (
	"heji-server/cmd"
	"heji-server/config"
	"os"
)

func main() {
	conf := config.Load("config.yml")
	cmd.Main(os.Args, conf)
}
