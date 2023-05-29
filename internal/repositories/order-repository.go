package repositories

import (
	"project/internal/models"
	"sync"
)

type OrderRepository interface {
	GetAllOrders() ([]*models.Order, error)
	GetOrderByID(id string) (*models.Order, error)
	UpdateOrder(order *models.Order) error
}

type InMemoryOrderRepository struct {
	db  map[string]*models.Order
	mux sync.Mutex
}

func NewInMemoryOrderRepository() *InMemoryOrderRepository {
	return &InMemoryOrderRepository{
		db: make(map[string]*models.Order),
	}
}

func (r *InMemoryOrderRepository) GetAllOrders() ([]*models.Order, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	orders := make([]*models.Order, 0, len(r.db))
	for _, o := range r.db {
		orders = append(orders, o)
	}
	return orders, nil
}

func (r *InMemoryOrderRepository) GetOrderByID(id string) (*models.Order, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	order, ok := r.db[id]
	if !ok {
		return nil, nil
	}
	return order, nil
}

func (r *InMemoryOrderRepository) UpdateOrder(order *models.Order) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	r.db[order.ID] = order
	return nil
}
