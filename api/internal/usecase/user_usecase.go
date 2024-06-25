package usecase
// "github.com/fs0414/go_hobby/internal/usecase"

import (
	"github.com/gin-gonic/gin"

	"github.com/fs0414/go_hobby/internal/adapter/repository/user"
	"github.com/fs0414/go_hobby/internal/infrastructure/model"
	"github.com/fs0414/go_hobby/internal/service/user"
)

type UserUseCase struct {
	repo repository.UserRepositoryImpl
	service service.UserServiceImpl
}

type SignInRequest struct {
	Email string `json:"email" binding:"required,min=1,max=255"`
}

func UserUseCaseFactory(
	repo repository.UserRepositoryImpl, 
	service service.UserServiceImpl,
) *UserUseCase {
	return &UserUseCase{repo: repo, service: service}
}

func (uc *UserUseCase) FetchUsers(c *gin.Context) {
	users, err := uc.repo.GetUsers()
	if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, users)
}

func (uc *UserUseCase) SignUp(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

	newUser, err := uc.repo.CreateUser(user)

	if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

	c.JSON(201, newUser)
}

func (uc *UserUseCase) SignIn(c *gin.Context) {
	var req SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

	user, err := uc.repo.FindByCredentials(req.Email)

	if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

	token, err := uc.service.PublishToken(user.ID)

	if err!= nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

	c.JSON(200, token)
}

