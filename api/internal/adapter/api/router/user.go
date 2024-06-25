package router

import (
    "github.com/gin-gonic/gin"

    "github.com/fs0414/go_hobby/internal/usecase"
	"github.com/fs0414/go_hobby/internal/adapter/repository/user"
	"github.com/fs0414/go_hobby/internal/service/user"
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