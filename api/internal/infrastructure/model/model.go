package model
// "github.com/fs0414/go_hobby/internal/infrastructure/model"

import (
	"gorm.io/gorm"
)

type UserRole string

const (
    RoleAdmin = "admin"
    RoleUser  = "user"
)

type User struct {
	gorm.Model
	Name string `gorm:"size:255";column:name`
	Email string `gorm:"type:varchar(100);unique_index;column:email"`
	Password string `gorm:"size:255";column:password`
    Role UserRole `gorm:"type:varchar(20);column:role"`
	Boards []Board `gorm:"foreignKey:UserID"`
}

type Board struct {
	gorm.Model
	Content string `gorm:"size:255;column:content"`
	UserID uint `gorm"column:user_id"`
}