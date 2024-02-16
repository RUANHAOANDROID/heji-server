package controller

import (
	"github.com/gin-gonic/gin"
	"heji-server/domain"
	"net/http"
)

type BookController struct {
	UseCase domain.BookUseCase
}

func (bc *BookController) CreateBook(c *gin.Context) {
	var book domain.Book
	err := c.ShouldBind(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.RespError("参数错误"))
		return
	}
	err = bc.UseCase.CreateBook(c, &book)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.RespError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, domain.RespSuccess(nil))
}
func (bc *BookController) BookList(c *gin.Context) {
	books, err := bc.UseCase.BookList(c, getUserId(c))
	if err != nil {
		c.JSON(http.StatusOK, domain.RespError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, domain.RespSuccess(books))
}

func (bc *BookController) DeleteBook(c *gin.Context) {

}
func (bc *BookController) JoinBook(c *gin.Context) {

}
func (bc *BookController) UpdateBook(c *gin.Context) {

}
func (bc *BookController) SharedBook(c *gin.Context) {

}
