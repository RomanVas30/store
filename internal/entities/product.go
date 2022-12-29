package entities

type Product struct {
	Id          int    `json:"-" db:"id"`
	Name        string `json:"name" db:"name"  binding:"required"`
	Cost        int    `json:"cost" db:"cost" binding:"required"`
	Count       int    `json:"count" db:"count" binding:"required"`
	Description string `json:"description" db:"description"`
}

type ShortProduct struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Cost int    `json:"cost" db:"cost"`
}

type AddProduct struct {
	OrderId      int `json:"order_id" db:"order_id"`
	UserId       int `json:"-" db:"user_id"`
	ProductId    int `json:"product_id" db:"product_id"`
	ProductCount int `json:"product_count" db:"product_count"`
}
