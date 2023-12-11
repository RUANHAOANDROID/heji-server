package controller

import (
	"github.com/gin-gonic/gin"
	"heji-server/domain"
	"net/http"
)

type LoginController struct {
	LoginUseCase domain.LoginUseCase
}

func (lc *LoginController) Login(c *gin.Context) {

	c.JSON(http.StatusOK, "login response ")
}
