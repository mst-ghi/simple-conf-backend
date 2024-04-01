package users

import (
	"simple-conf/core"
	"simple-conf/database/models"
	"simple-conf/database/repositories"
)

type UsersServiceInterface interface {
	FindAll(exceptUserId string) []models.User
	FindOne(id string) (models.User, core.Error)
}

type UsersService struct {
	repository repositories.UserRepositoryInterface
}

func NewUsersService() *UsersService {
	return &UsersService{
		repository: repositories.NewUserRepository(),
	}
}

func (service *UsersService) FindAll(exceptUserId string) []models.User {
	return service.repository.FindAll(exceptUserId)
}

func (service *UsersService) FindOne(id string) (models.User, core.Error) {
	user := service.repository.FindByID(id)

	if user.ID == "" {
		return user, core.Error{"reason": "User not found"}
	}

	return user, nil
}
