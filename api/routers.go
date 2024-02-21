package api

import (
	"github.com/gin-gonic/gin"
	"heji-server/api/controller"
	"heji-server/config"
	"heji-server/domain"
	"heji-server/internal/get"
	"heji-server/mongo"
	"heji-server/repository"
	"heji-server/usecase"
	"time"
)

// registerRoutes configures the available web server routes.
func registerRoutes(conf *config.Config, db mongo.Database) {
	timeout := get.Config().Mongo.TimeoutMax
	WebSocket(APIv1, conf)
	NewUserRouter(db, timeout, APIv1)
	NewBookRouter(db, timeout, APIv1)
}
func NewUserRouter(db mongo.Database, timeout time.Duration, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollUser)
	lc := &controller.UserController{
		UserUseCase: usecase.NewLoginUseCase(ur, timeout),
	}
	group.POST("/Register", lc.Register)
	group.POST("/Login", lc.Login)
}
func NewBookRouter(db mongo.Database, timeout time.Duration, group *gin.RouterGroup) {
	br := repository.NewBookRepository(db, domain.CollBook)
	sr := repository.NewSharedRepository(db, domain.CollShared)
	bc := &controller.BookController{
		UseCase: usecase.NewBookUseCase(br, sr, timeout),
	}
	group.POST("/CreateBook", bc.CreateBook)
	group.POST("/BookList", bc.BookList)
	group.POST("/DeleteBook", bc.DeleteBook)
	group.POST("/UpdateBook", bc.UpdateBook)
	group.POST("/SharedBook", bc.SharedBook)
}
