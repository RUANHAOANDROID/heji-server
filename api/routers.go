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
	WebSocket(APIv1, conf)
	NewUserRouter(db, get.Config().Mongo.TimeoutMax, APIv1)
}
func NewUserRouter(db mongo.Database, timeout time.Duration, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollUser)
	lc := &controller.UserController{
		UserUseCase: usecase.NewLoginUseCase(ur, timeout),
	}
	group.POST("/Register", lc.Register)
	group.POST("/Login", lc.Login)
}
