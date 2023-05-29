package repositories

import (
	"project/internal/models"
	"sync"
)

type ProductRepository interface {
	GetAllProducts() ([]*models.Product, error)
	GetProductByID(id string) (*models.Product, error)
	UpdateProduct(product *models.Product) error
}

type InMemoryProductRepository struct {
	db  map[string]*models.Product
	mux sync.Mutex
}

func NewInMemoryProductRepository() *InMemoryProductRepository {
	return &InMemoryProductRepository{
		db: make(map[string]*models.Product),
	}
}

func (r *InMemoryProductRepository) GetAllProducts() ([]*models.Product, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	products := make([]*models.Product, 0, len(r.db))
	for _, p := range r.db {
		products = append(products, p)
	}
	return products, nil
}

func (r *InMemoryProductRepository) GetProductByID(id string) (*models.Product, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	product, ok := r.db[id]
	if !ok {
		return nil, nil
	}
	return product, nil
}

func (r *InMemoryProductRepository) UpdateProduct(product *models.Product) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	r.db[product.ID] = product
	return nil
}
