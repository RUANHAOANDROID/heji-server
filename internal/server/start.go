package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"heji-server/internal/config"
	"time"
)

func Start(ctx context.Context, conf *config.Config) {
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
		}
	}()
	start := time.Now()
	log.Info(start)
	if conf.HttpMode() != "" {
		gin.SetMode(conf.HttpMode())
	}
	router := gin.New()

	if err := router.SetTrustedProxies(conf.TrustedProxies()); err != nil {
		log.Warnf("server: %s", err)
	}
	router.Use(Recovery(), Security(conf))
	APIv1 = router.Group(config.ApiUri)
	registerRoutes(router, conf)
}
