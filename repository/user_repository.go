package repository

import (
	"github.com/wisnuragaprawida/golang-rest-api/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Login(User *model.User) ([]model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func (u *userRepository) Login(User *model.User) ([]model.User, error) {
	var users []model.User
	err := u.db.Where("email = ? AND password= ?", User.Email, User.Password).First(&users).Error

	if err != nil {
		return nil, err
	}
	return users, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	repo := new(userRepository)
	repo.db = db
	return repo
}
