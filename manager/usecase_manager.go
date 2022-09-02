package manager

import "github.com/wisnuragaprawida/golang-rest-api/usecase"

type UseCaseManager interface {
	CreateProductUseCase() usecase.CreateProductUseCase
	ListProductUseCase() usecase.ListProductUseCase
	LoginUserUseCase() usecase.LoginUserUseCase
}

type useCaseManager struct {
	repoManager RepositoryManager
}

func (u *useCaseManager) CreateProductUseCase() usecase.CreateProductUseCase {
	return usecase.NewCreateProductUseCase(u.repoManager.ProductRepo())
}

func (u *useCaseManager) ListProductUseCase() usecase.ListProductUseCase {
	return usecase.NewListProductUseCase(u.repoManager.ProductRepo())
}

func (u *useCaseManager) LoginUserUseCase() usecase.LoginUserUseCase {
	return usecase.NewLoginUserUseCase(u.repoManager.UserRepo())
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}
