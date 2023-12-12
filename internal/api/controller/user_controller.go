package controller

import (
	"github.com/gin-gonic/gin"
	"heji-server/domain"
	"net/http"
)

type UserController struct {
	LoginUseCase domain.LoginUseCase
}

func (uc *UserController) Login(c *gin.Context) {

	c.JSON(http.StatusOK, "login response ")
}
