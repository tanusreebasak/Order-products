package services

import (
	"project/internal/models"
	"project/internal/repositories"
)

type ProductService interface {
	GetAllProducts() ([]*models.Product, error)
	GetProductByID(id string) (*models.Product, error)
}

type ProductServiceImpl struct {
	productRepo repositories.ProductRepository
}

func NewProductService(productRepo repositories.ProductRepository) *ProductServiceImpl {
	return &ProductServiceImpl{
		productRepo: productRepo,
	}
}

func (s *ProductServiceImpl) GetAllProducts() ([]*models.Product, error) {
	return s.productRepo.GetAllProducts()
}

func (s *ProductServiceImpl) GetProductByID(id string) (*models.Product, error) {
	return s.productRepo.GetProductByID(id)
}
