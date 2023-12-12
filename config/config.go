package config

import (
	"heji-server/pkg"
	"time"
)

// Config app config
type Config struct {
	App     App
	Mongo   Mongo
	Jwt     Jwt
	Options Options
}
type App struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type Mongo struct {
	Host       string        `yaml:"host"`
	Port       string        `yaml:"port"`
	Database   string        `yaml:"database"`
	Username   string        `yaml:"username"`
	Password   string        `yaml:"password"`
	TimeoutMax time.Duration `yaml:"timeoutMax"`
}
type Jwt struct {
	Secret         string        `yaml:"secret"`
	ExpirationTime time.Duration `yaml:"expirationTime"`
}
type Options struct {
}

func Load(path string) *Config {
	// Initialize options from config file and CLI context.
	var config Config
	pkg.LoadYml(path, &config)
	return &config
}
func Get() {

}
