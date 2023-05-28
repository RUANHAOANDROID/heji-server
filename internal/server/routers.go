package server

import (
	"github.com/gin-gonic/gin"
	"heji-server/internal/api"
	"heji-server/internal/config"
)

var APIv1 *gin.RouterGroup

// registerRoutes configures the available web server routes.
func registerRoutes(router *gin.Engine, conf *config.Config) {
	api.HandlerHoldWS(APIv1, conf)
	api.Login(APIv1)
}
