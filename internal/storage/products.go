package storage

import (
	"github.com/RomanVas30/store/external/dbr_extensions"
	"github.com/RomanVas30/store/internal/entities"
	"github.com/gocraft/dbr"
)

type Products interface {
	CreateProduct(product *entities.Product) error
	GetProducts() (*[]entities.ShortProduct, error)
	GetProductById(productId int) (*entities.Product, error)
}

type ProductsStorage struct {
	db *dbr.Connection
}

func NewProductsStorage(db *dbr.Connection) *ProductsStorage {
	return &ProductsStorage{
		db: db,
	}
}

func (ps *ProductsStorage) CreateProduct(product *entities.Product) error {
	newSession := ps.db.NewSession(nil)

	sessionError := dbr_extensions.CreateTx(newSession, func(runner dbr.SessionRunner) error {
		err := runner.InsertInto("product").
			Pair("name", product.Name).
			Pair("cost", product.Cost).
			Pair("count", product.Count).
			Pair("description", product.Description).
			Returning("id").
			Load(&product.Id)
		if err != nil {
			return err
		}

		return nil
	})
	if sessionError != nil {
		return sessionError
	}

	return nil
}

func (ps *ProductsStorage) GetProducts() (*[]entities.ShortProduct, error) {
	newSession := ps.db.NewSession(nil)

	var products []entities.ShortProduct

	sessionError := dbr_extensions.CreateTx(newSession, func(runner dbr.SessionRunner) error {
		_, err := runner.Select("id", "name", "cost").
			From("product").
			Load(&products)
		if err != nil {
			return err
		}
		return nil
	})
	if sessionError != nil {
		return nil, sessionError
	}

	return &products, nil
}

func (ps *ProductsStorage) GetProductById(productId int) (*entities.Product, error) {
	newSession := ps.db.NewSession(nil)

	var productInfo entities.Product

	sessionError := dbr_extensions.CreateTx(newSession, func(runner dbr.SessionRunner) error {
		_, err := runner.Select("id", "name", "cost", "count", "description").
			From("product").
			Where(dbr.Eq("id", productId)).
			Load(&productInfo)
		if err != nil {
			return err
		}
		return nil
	})
	if sessionError != nil {
		return nil, sessionError
	}

	return &productInfo, nil
}
