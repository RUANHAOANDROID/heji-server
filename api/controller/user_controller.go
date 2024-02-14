package controller

import (
	"github.com/gin-gonic/gin"
	"heji-server/domain"
	"net/http"
)

type UserController struct {
	UserUseCase domain.UserUseCase
}

func (uc *UserController) Login(c *gin.Context) {
	var request domain.LoginRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.RespError("Login request format error!"))
		return
	}
	token, err := uc.UserUseCase.Login(c, &request)

	c.JSON(http.StatusOK, domain.RespSuccess(token))
}
func (uc *UserController) Register(c *gin.Context) {
	var user domain.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.RespError(err.Error()))
		return
	}
	err = uc.UserUseCase.Register(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.RespError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, domain.RespSuccess(""))
}
