package api

import (
	"github.com/gin-gonic/gin"
	"heji-server/config"
	"heji-server/pkg"
)

var log = pkg.Log

var APIv1 *gin.RouterGroup

func Setup(conf *config.Config) {
	router := gin.Default()
	trustedProxies := []string{
		"127.0.0.1",
		//config.Localhost,
	}
	if err := router.SetTrustedProxies(trustedProxies); err != nil {
		log.Warnf("server: %s", err)
	}
	//开启重定向
	router.RedirectTrailingSlash = true
	APIv1 = router.Group(config.ApiUri)
	registerRoutes(conf)

	router.Run(":8888")
}

// registerRoutes configures the available web server routes.
func registerRoutes(conf *config.Config) {
	WebSocket(APIv1, conf)
}
