package usecase

import (
	"github.com/wisnuragaprawida/golang-rest-api/model"
	"github.com/wisnuragaprawida/golang-rest-api/repository"
)

type LoginUserUseCase interface {
	LoginUser(newProduct *model.User) ([]model.User, error)
}

type loginUserUseCase struct {
	repo repository.UserRepository
}

func (l *loginUserUseCase) LoginUser(newProduct *model.User) ([]model.User, error) {
	return l.repo.Login(newProduct)
}

func NewLoginUserUseCase(repo repository.UserRepository) LoginUserUseCase {
	return &loginUserUseCase{repo: repo}
}
