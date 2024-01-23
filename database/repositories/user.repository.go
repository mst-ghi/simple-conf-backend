package repositories

import (
	"video-conf/database"
	"video-conf/database/models"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	Connection() *gorm.DB
	Create(user models.User) models.User
	FindByEmail(email string) models.User
	FindByID(id string) models.User
	FindAll() []models.User
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		DB: database.Connection(),
	}
}

func (self *UserRepository) Connection() *gorm.DB {
	return self.DB
}

func (self *UserRepository) Create(user models.User) models.User {
	var newUser models.User
	self.DB.Create(&user).Scan(&newUser)
	return newUser
}

func (self *UserRepository) FindByEmail(email string) models.User {
	var user models.User
	self.DB.First(&user, "email = ?", email)
	return user
}

func (self *UserRepository) FindByID(id string) models.User {
	var user models.User
	self.DB.First(&user, "id = ?", id)
	return user
}

func (self *UserRepository) FindAll() []models.User {
	var users []models.User
	self.DB.Table("users").Find(&users)
	return users
}
