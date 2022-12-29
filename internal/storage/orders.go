package storage

import (
	"fmt"
	"github.com/RomanVas30/store/external/dbr_extensions"
	"github.com/RomanVas30/store/internal/entities"
	"github.com/gocraft/dbr"
)

type Orders interface {
	CreateOrder(product *entities.Order) error
	GetOrders(userId int) (*[]entities.Order, error)
	GetOrderById(orderId int, userId int) (*entities.OrderWithProducts, error)
	OrderPayment(orderId int, userId int) error
	AddProductToOrder(addProduct *entities.AddProduct) error
}

type OrdersStorage struct {
	db *dbr.Connection
}

func NewOrdersStorage(db *dbr.Connection) *OrdersStorage {
	return &OrdersStorage{
		db: db,
	}
}

func (os *OrdersStorage) CreateOrder(order *entities.Order) error {
	newSession := os.db.NewSession(nil)

	sessionError := dbr_extensions.CreateTx(newSession, func(runner dbr.SessionRunner) error {
		err := runner.InsertInto("order_table").
			Pair("name", order.Name).
			Pair("user_id", order.UserId).
			Returning("id").
			Load(&order.Id)
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

func (os *OrdersStorage) GetOrders(userId int) (*[]entities.Order, error) {
	newSession := os.db.NewSession(nil)

	var orders []entities.Order

	sessionError := dbr_extensions.CreateTx(newSession, func(runner dbr.SessionRunner) error {
		_, err := runner.Select("id", "name", "status").
			From("order_table").
			Where(dbr.Eq("user_id", userId)).
			Load(&orders)
		if err != nil {
			return err
		}
		return nil
	})
	if sessionError != nil {
		return nil, sessionError
	}

	return &orders, nil
}

func (os *OrdersStorage) GetOrderById(orderId int, userId int) (*entities.OrderWithProducts, error) {
	newSession := os.db.NewSession(nil)

	var orderInfo entities.OrderWithProducts

	sessionError := dbr_extensions.CreateTx(newSession, func(runner dbr.SessionRunner) error {
		rowsCount, err := runner.Select("name", "status").
			From("order_table").
			Where(dbr.Eq("id", orderId)).
			Where(dbr.Eq("user_id", userId)).
			Load(&orderInfo)
		if err != nil {
			return err
		}
		if rowsCount != 1 {
			return fmt.Errorf("order with id = %d not found", orderId)
		}

		_, err = runner.Select("opt.product_id", "pr.name as product_name", "opt.product_count").
			From(dbr.I("order_products_table").As("opt")).
			LeftJoin(dbr.I("product").As("pr"), dbr.Expr("opt.product_id = pr.id")).
			Where(dbr.Eq("order_id", orderId)).
			Load(&orderInfo.Products)
		if err != nil {
			return err
		}

		return nil
	})
	if sessionError != nil {
		return nil, sessionError
	}

	return &orderInfo, nil
}

func (os *OrdersStorage) OrderPayment(orderId int, userId int) error {
	newSession := os.db.NewSession(nil)

	sessionError := dbr_extensions.CreateTx(newSession, func(runner dbr.SessionRunner) error {
		result, err := runner.Update("order_table").
			Set("status", "IS_PAID").
			Where(dbr.Eq("id", orderId)).
			Where(dbr.Eq("user_id", userId)).
			Where(dbr.Eq("status", "IS_NOT_PAID")).
			Exec()
		if err != nil {
			return err
		}
		if rows, _ := result.RowsAffected(); rows != 1 {
			return fmt.Errorf("order payment error with id = %d", orderId)
		}

		return nil
	})
	if sessionError != nil {
		return sessionError
	}

	return nil
}

func (os *OrdersStorage) AddProductToOrder(addProduct *entities.AddProduct) error {
	newSession := os.db.NewSession(nil)

	sessionError := dbr_extensions.CreateTx(newSession, func(runner dbr.SessionRunner) error {
		var orderInfo entities.Order
		rowsCount, err := runner.Select("name").
			From("order_table").
			Where(dbr.Eq("id", addProduct.OrderId)).
			Where(dbr.Eq("user_id", addProduct.UserId)).
			Load(&orderInfo)
		if err != nil {
			return err
		}
		if rowsCount != 1 {
			return fmt.Errorf("order with id = %d not found", addProduct.OrderId)
		}

		var productInfo entities.Product
		rowsCount, err = runner.Select("count").
			From("product").
			Where(dbr.Eq("id", addProduct.ProductId)).
			Load(&productInfo)
		if err != nil {
			return err
		}
		if rowsCount != 1 {
			return fmt.Errorf("product with id = %d not found", addProduct.ProductId)
		}

		if productInfo.Count < addProduct.ProductCount {
			return fmt.Errorf("product with id = %d insufficient quantity", addProduct.ProductId)
		}

		_, err = runner.InsertInto("order_products_table").
			Pair("order_id", addProduct.OrderId).
			Pair("product_id", addProduct.ProductId).
			Pair("product_count", addProduct.ProductCount).
			Exec()
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
