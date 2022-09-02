package usecase

import (
	"github.com/wisnuragaprawida/golang-rest-api/model"
	"github.com/wisnuragaprawida/golang-rest-api/repository"
)

type CreateProductUseCase interface {
	CreateProduct(newProduct *model.Product) error
}

type createProductUseCase struct {
	repo repository.ProductRepository
}

func (c *createProductUseCase) CreateProduct(newProduct *model.Product) error {
	return c.repo.Add(newProduct)
}

func NewCreateProductUseCase(repo repository.ProductRepository) CreateProductUseCase {
	return &createProductUseCase{
		repo: repo,
	}
}
