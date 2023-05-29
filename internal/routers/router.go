package routers

import (
	"net/http"
	"project/internal/controllers"
	"project/internal/repositories"
	"project/internal/services"
)

func SetupRouter() http.Handler {
	productRepo := repositories.NewInMemoryProductRepository()
	orderRepo := repositories.NewInMemoryOrderRepository()

	productService := services.NewProductService(productRepo)
	orderService := services.NewOrderService(orderRepo, productRepo, 10, 3)

	productController := controllers.NewProductController(productService)
	orderController := controllers.NewOrderController(orderService)

	router := http.NewServeMux()

	router.HandleFunc("/products", productController.GetAllProducts)
	router.HandleFunc("/products/id", productController.GetProductByID)

	router.HandleFunc("/orders", orderController.GetAllOrders)
	router.HandleFunc("/orders/id", orderController.GetOrderByID)
	router.HandleFunc("/orders/place", orderController.PlaceOrder)
	router.HandleFunc("/orders/update-status", orderController.UpdateOrderStatus)

	return router
}
