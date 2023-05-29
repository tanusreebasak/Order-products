package services

import (
	"project/internal/models"
	"project/internal/repositories"
)

type OrderService interface {
	GetAllOrders() ([]*models.Order, error)
	GetOrderByID(id string) (*models.Order, error)
	PlaceOrder(order *models.Order) error
	UpdateOrderStatus(orderID string, status string) error
}

type OrderServiceImpl struct {
	orderRepo    repositories.OrderRepository
	productRepo  repositories.ProductRepository
	maxQuantity  int
	premiumCount int
}

func NewOrderService(orderRepo repositories.OrderRepository, productRepo repositories.ProductRepository, maxQuantity, premiumCount int) *OrderServiceImpl {
	return &OrderServiceImpl{
		orderRepo:    orderRepo,
		productRepo:  productRepo,
		maxQuantity:  maxQuantity,
		premiumCount: premiumCount,
	}
}

func (s *OrderServiceImpl) GetAllOrders() ([]*models.Order, error) {
	return s.orderRepo.GetAllOrders()
}

func (s *OrderServiceImpl) GetOrderByID(id string) (*models.Order, error) {
	return s.orderRepo.GetOrderByID(id)
}

func (s *OrderServiceImpl) PlaceOrder(order *models.Order) error {
	product, err := s.productRepo.GetProductByID(order.ProductID)
	if err != nil {
		return err
	}
	if product == nil {
		return ErrProductNotFound
	}
	if order.Quantity > s.maxQuantity {
		return ErrExceededMaxQuantity
	}

	if product.Category == "Premium" {
		s.premiumCount++
		if s.premiumCount >= 3 {
			order.DiscountApplied = true
		}
	}

	order.OrderValue = float64(order.Quantity) * product.Price
	if order.DiscountApplied {
		order.OrderValue *= 0.9
	}

	err = s.orderRepo.UpdateOrder(order)
	if err != nil {
		return err
	}

	product.Quantity -= order.Quantity
	err = s.productRepo.UpdateProduct(product)
	if err != nil {
		return err
	}

	return nil
}

func (s *OrderServiceImpl) UpdateOrderStatus(orderID string, status string) error {
	order, err := s.orderRepo.GetOrderByID(orderID)
	if err != nil {
		return err
	}
	if order == nil {
		return ErrOrderNotFound
	}

	order.Status = status
	if status == "Dispatched" {
		order.DispatchDate = getCurrentDate()
	}

	return s.orderRepo.UpdateOrder(order)
}
