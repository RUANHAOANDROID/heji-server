package controller

import (
	"github.com/gin-gonic/gin"
	"heji-server/domain"
	"net/http"
)

type BookController struct {
	UseCase domain.BookUseCase
}

func (bc *BookController) Create(c *gin.Context) {
	c.JSON(http.StatusOK, "response")
}
