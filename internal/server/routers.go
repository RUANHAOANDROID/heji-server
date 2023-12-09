package server

import (
	"github.com/gin-gonic/gin"
	"heji-server/config"
	"heji-server/internal/api"
)

var APIv1 *gin.RouterGroup

// registerRoutes configures the available web server routes.
func registerRoutes(router *gin.Engine, conf *config.Config) {
	//开启重定向
	router.RedirectTrailingSlash = true
	registerStaticRoutes(router, conf)
	api.WebSocket(APIv1, conf)
	api.Login(APIv1)
	log.Debug("路由注册完毕")
}
