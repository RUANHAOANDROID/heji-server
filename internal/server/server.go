package server

import (
	"github.com/gin-gonic/gin"
	config2 "heji-server/config"
	"heji-server/pkg"
)

var log = pkg.Log

func Start(conf *config2.Config) {
	router := gin.Default()
	trustedProxies := []string{
		"127.0.0.1",
		//config.Localhost,
	}
	if err := router.SetTrustedProxies(trustedProxies); err != nil {
		log.Warnf("server: %s", err)
	}
	router.Use(Recovery(), Security(conf))
	APIv1 = router.Group(config2.ApiUri)
	registerRoutes(router, conf)

	router.Run(":8888")
}
