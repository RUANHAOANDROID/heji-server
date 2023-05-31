package server

import (
	"github.com/gin-gonic/gin"
	"heji-server/internal/config"
	"net/http"
)

// registerStaticRoutes configures serving static assets and templates.
func registerStaticRoutes(router *gin.Engine, conf *config.Config) {
	router.NoRoute(func(c *gin.Context) {
		switch c.NegotiateFormat(gin.MIMEHTML, gin.MIMEJSON) {
		case gin.MIMEJSON:
			c.JSON(http.StatusNotFound, gin.H{"error": "i18n.Msg(i18n.ErrNotFound)"})
		default:
			values := gin.H{
				"signUp": gin.H{"message": "config.MsgSponsor", "url": " config.SignUpURL"},
				"config": "conf.ClientPublic()",
				"error":  "ErrNotFound",
				"code":   http.StatusNotFound,
			}
			c.HTML(http.StatusNotFound, "404.gohtml", values)
		}
	})
}
