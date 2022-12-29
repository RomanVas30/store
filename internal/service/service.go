package service

import (
	"github.com/RomanVas30/store/internal/storage"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Service struct {
	Authorization
	Staff
	OrgUnits
	Products
	Orders
}

func NewService(repos *storage.Storage) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Staff:         NewStaffService(repos.Staff),
		OrgUnits:      NewOrgUnitsService(repos.OrgUnits),
		Products:      NewProductsService(repos.Products),
		Orders:        NewOrdersService(repos.Orders),
	}
}
