package service

import (
	"github.com/RomanVas30/store/internal/entities"
	"github.com/RomanVas30/store/internal/storage"
)

type Products interface {
	CreateProduct(product *entities.Product) error
	GetProducts() (*[]entities.ShortProduct, error)
	GetProductById(productId int) (*entities.Product, error)
}

type ProductsService struct {
	repo storage.Products
}

func NewProductsService(repo storage.Products) *ProductsService {
	return &ProductsService{repo: repo}
}

func (s *ProductsService) CreateProduct(product *entities.Product) error {
	return s.repo.CreateProduct(product)
}

func (s *ProductsService) GetProducts() (*[]entities.ShortProduct, error) {
	return s.repo.GetProducts()
}

func (s *ProductsService) GetProductById(productId int) (*entities.Product, error) {
	return s.repo.GetProductById(productId)
}
