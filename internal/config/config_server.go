package config

const (
	HttpModeProd  = "release"
	HttpModeDebug = "debug"
)

// HttpMode returns the server mode.
func (c *Config) HttpMode() string {
	if c.Prod() {
		return HttpModeProd
	} else {
		return HttpModeDebug
	}
}
