package controller

import (
	"github.com/gin-gonic/gin"
	"heji-server/domain"
	"net/http"
)

type BookController struct {
	BookUseCase domain.BookUseCase
}

func (bc *BookController) Create(c *gin.Context) {
	//var request domain.Book
	c.JSON(http.StatusOK, "response")
}
