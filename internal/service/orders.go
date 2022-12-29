package service

import (
	"github.com/RomanVas30/store/internal/entities"
	"github.com/RomanVas30/store/internal/storage"
)

type Orders interface {
	CreateOrder(product *entities.Order) error
	GetOrders(userId int) (*[]entities.Order, error)
	GetOrderById(orderId int, userId int) (*entities.OrderWithProducts, error)
	OrderPayment(orderId int, userId int) error
	AddProductToOrder(addProduct *entities.AddProduct) error
}

type OrdersService struct {
	repo storage.Orders
}

func NewOrdersService(repo storage.Orders) *OrdersService {
	return &OrdersService{repo: repo}
}

func (s *OrdersService) CreateOrder(order *entities.Order) error {
	return s.repo.CreateOrder(order)
}

func (s *OrdersService) GetOrders(userId int) (*[]entities.Order, error) {
	return s.repo.GetOrders(userId)
}

func (s *OrdersService) GetOrderById(orderId int, userId int) (*entities.OrderWithProducts, error) {
	return s.repo.GetOrderById(orderId, userId)
}

func (s *OrdersService) OrderPayment(orderId int, userId int) error {
	return s.repo.OrderPayment(orderId, userId)
}

func (s *OrdersService) AddProductToOrder(addProduct *entities.AddProduct) error {
	return s.repo.AddProductToOrder(addProduct)
}
