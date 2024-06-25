package repository

import (
	"github.com/fs0414/go_hobby/internal/infrastructure/model"
)

type UserRepositoryImpl interface {
	GetUsers() ([]model.User, error)
	CreateUser(user model.User) (model.User, error)
	FindByCredentials(email string) (model.User, error)
}