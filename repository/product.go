package repository

import (
	"go-jwt/model"

	"github.com/jinzhu/gorm"
)

//ProductRepository --> Interface to ProductRepository
type ProductRepository interface {
	GetProduct(int) (model.Product, error)
	GetAllProducts() ([]model.Product, error)
}

type productRepository struct {
	connection *gorm.DB
}

//NewProductRepository --> returns new product repository
func NewProductRepository() ProductRepository {
	return &productRepository{
		connection: DB(),
	}
}

func (db *productRepository) GetProduct(id int) (product model.Product, err error) {
	return product, db.connection.First(&product, id).Error
}

func (db *productRepository) GetAllProducts() (products []model.Product, err error) {
	return products, db.connection.Find(&products).Error
}