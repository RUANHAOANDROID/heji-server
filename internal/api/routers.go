package api

import (
	"github.com/gin-gonic/gin"
	"heji-server/config"
	"heji-server/domain"
	"heji-server/internal/api/controller"
	"heji-server/internal/get"
	"heji-server/mongo"
	"heji-server/repository"
	"heji-server/usecase"
	"time"
)

// registerRoutes configures the available web server routes.
func registerRoutes(conf *config.Config, db mongo.Database) {
	WebSocket(APIv1, conf)
	NewLoginRouter(db, get.Config().Mongo.TimeoutMax, APIv1)
}
func NewLoginRouter(db mongo.Database, timeout time.Duration, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollUser)
	lc := &controller.LoginController{
		LoginUseCase: usecase.NewLoginUseCase(ur, timeout),
	}
	group.POST("/login", lc.Login)
}
