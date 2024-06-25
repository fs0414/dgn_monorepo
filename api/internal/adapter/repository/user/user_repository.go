package repository
// "github.com/fs0414/go_hobby/internal/adapter/repository"

import (
	"github.com/fs0414/go_hobby/internal/infrastructure/database"
	"github.com/fs0414/go_hobby/internal/infrastructure/model"
)

type UserRepository struct {}

func UserRepositoryFactory() UserRepositoryImpl {
	return &UserRepository{}
}

func (repo *UserRepository) GetUsers() ([]model.User, error) {
	db := database.GetDb()
	var users []model.User
    result := db.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (repo *UserRepository) CreateUser(user model.User) (model.User, error) {
	db := database.GetDb()
	result := db.Create(&user)
	if result.Error != nil {
		return model.User{}, result.Error
	}

	return user, nil
}

func (repo *UserRepository) FindByCredentials(email string) (model.User, error) {
	db := database.GetDb()
    var user model.User
    
	if err := db.Where("email =?", email).First(&user).Error; err != nil {
		return model.User{}, err
    }
	
    return user, nil
}