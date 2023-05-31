package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(router *gin.RouterGroup) {
	router.GET("/user/login", func(c *gin.Context) {

		c.JSONP(http.StatusOK, "username = test")
	})
}
