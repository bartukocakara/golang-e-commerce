package repository

import (
	"go-jwt/model"

	"github.com/jinzhu/gorm"
)

//UserRepository -> User CRUD
type UserRepository interface {
	AddUser(model.User)(model.User, error)
	GetUser(int) (model.User, error)
	GetByEmail(string) (model.User, error)
	GetAllUser() ([]model.User, error)

}

type userRepository struct {
	connection *gorm.DB
}

//NewUserRepository --> returns new user repository
func NewUserRepository() UserRepository {
	return &userRepository{
		connection: DB(),
	}
}

func (db *userRepository) GetUser(id int) (user model.User, err error) {
	return user, db.connection.First(&user, id).Error
}

func (db *userRepository) AddUser(user model.User) (model.User, error) {
	return user, db.connection.Create(&user).Error
}

func (db *userRepository) GetByEmail(email string) (user model.User, err error) {
	return user, db.connection.First(&user, "email=?", email).Error
}

func (db *userRepository) GetAllUser() (users []model.User, err error) {
	return users, db.connection.Find(&users).Error
}