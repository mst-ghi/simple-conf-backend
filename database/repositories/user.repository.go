package repositories

import (
	"fmt"
	"video-conf/database"
	"video-conf/database/models"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	Connection() *gorm.DB
	Create(user models.User) models.User
	FindByEmail(email string) models.User
	FindByID(id string) models.User
	FindAll(exceptUserId string) []models.User
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		DB: database.Connection(),
	}
}

func (repo *UserRepository) Connection() *gorm.DB {
	return repo.DB
}

func (repo *UserRepository) Create(user models.User) models.User {
	var newUser models.User
	repo.DB.Create(&user).Scan(&newUser)
	return newUser
}

func (repo *UserRepository) FindByEmail(email string) models.User {
	var user models.User
	repo.DB.First(&user, "email = ?", email)
	return user
}

func (repo *UserRepository) FindByID(id string) models.User {
	var user models.User
	repo.DB.First(&user, "id = ?", id)
	return user
}

func (repo *UserRepository) FindAll(exceptUserId string) []models.User {
	fmt.Println("exceptUserId", exceptUserId)
	var users []models.User
	repo.DB.Table("users").Not("id = ?", exceptUserId).Find(&users)
	return users
}
