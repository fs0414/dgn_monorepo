package router

import (
	"github.com/gin-gonic/gin"

	repository "github.com/fs0414/internal/adapter/repository/user"
	service "github.com/fs0414/internal/service/user"
	"github.com/fs0414/internal/usecase"
)

func SetUserRoutes(rg *gin.RouterGroup) {
	uc := usecase.UserUseCaseFactory(
		repository.UserRepositoryFactory(),
		service.UserServiceFactory(),
	)

	rg.GET("/users", uc.FetchUsers)
	rg.POST("/signup", uc.SignUp)
	rg.POST("/signin", uc.SignIn)
}
