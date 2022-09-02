package manager

import (
	"github.com/wisnuragaprawida/golang-rest-api/repository"
)

type RepositoryManager interface {
	// ProductRepo Disini kumpulan semua repo dalam 1 project yang dibuat
	ProductRepo() repository.ProductRepository
	UserRepo() repository.UserRepository
}

type repositoryManager struct {
	infra Infra
}

func (r *repositoryManager) ProductRepo() repository.ProductRepository {
	return repository.NewProductRepository(r.infra.SqlDb())
}
func (r *repositoryManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.SqlDb())
}

func NewRepositoryManager(infra Infra) RepositoryManager {
	return &repositoryManager{
		infra: infra,
	}
}
