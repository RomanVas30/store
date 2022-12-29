package entities

type Order struct {
	Id     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name"  binding:"required"`
	UserId int    `json:"-" db:"user_id"`
	Status string `json:"status" db:"status"`
}

type OrderWithProducts struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" db:"name" binding:"required"`
	Status   string `json:"status" db:"status"`
	Products []struct {
		Id    int    `json:"product_id" db:"product_id" binding:"required"`
		Name  string `json:"product_name" db:"product_name" binding:"required"`
		Count string `json:"product_count" db:"product_count" binding:"required"`
	} `json:"products"`
}
