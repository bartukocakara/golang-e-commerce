package model

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Status   string  `json:"status"`
}

func (Product) TableName() string {
	return "products"
}