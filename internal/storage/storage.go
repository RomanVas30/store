package storage

import "github.com/gocraft/dbr"

type Storage struct {
	Authorization
	Staff
	OrgUnits
	Products
	Orders
}

func NewStorage(db *dbr.Connection) *Storage {
	return &Storage{
		Authorization: NewAuth(db),
		Staff:         NewStaffStorage(db),
		OrgUnits:      NewOrgUnitsStorage(db),
		Products:      NewProductsStorage(db),
		Orders:        NewOrdersStorage(db),
	}
}
