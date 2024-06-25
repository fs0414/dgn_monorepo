package service
// "github.com/fs0414/go_hobby/internal/service/user"

import (
	"fmt"
	"time"
	"os"

	"github.com/golang-jwt/jwt/v5"

	"github.com/fs0414/go_hobby/internal/infrastructure/database"
)

type UserService struct {}

func UserServiceFactory() UserServiceImpl {
    return &UserService{}
}

func (service *UserService) PublishToken(userId uint) (string, error) {
	fmt.Println("publish token")

	claims := jwt.MapClaims{
        "user_id": userId,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    }

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := locadEnvByJwtSecretKey()

    tokenString, err := token.SignedString([]byte(secretKey))

	return tokenString, err
}

func locadEnvByJwtSecretKey() string {
	database.LoadEnv()
	secretKey := os.Getenv("JWT_SECRET_KEY")
	return secretKey
}