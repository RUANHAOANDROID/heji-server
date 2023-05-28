package config

type Options struct {
	AppName        string   `yaml:"AppName" json:"AppName" flag:"AppName"`
	Prod           bool     `yaml:"Prod" json:"Prod" flag:"prod"`
	Debug          bool     `yaml:"Debug" json:"Debug" flag:"debug"`
	TrustedProxies []string `yaml:"TrustedProxies" json:"-" flag:"trusted-proxy"`
}
