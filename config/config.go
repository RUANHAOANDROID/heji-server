package config

import "heji-server/pkg"

// Config app config
type Config struct {
	App     App
	Mongodb Mongodb
	Jwt     Jwt
	Options Options
}
type App struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type Mongodb struct {
	Host     string `yaml:"host"`
	port     int    `yaml:"port"`
	database string `yaml:"database"`
	username string `yaml:"username"`
	password string `yaml:"password"`
}
type Jwt struct {
	Secret     string
	RememberMe int
}
type Options struct {
}

func Load() *Config {
	// Initialize options from config file and CLI context.
	var config Config
	pkg.LoadYml("./config", config)
	return &config
}
