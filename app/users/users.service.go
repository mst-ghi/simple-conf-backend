package users

import (
	"video-conf/core"
	"video-conf/database/models"
	"video-conf/database/repositories"
)

type UsersServiceInterface interface {
	FindAll() []models.User
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

func (service *UsersService) FindAll() []models.User {
	return service.repository.FindAll()
}

func (service *UsersService) FindOne(id string) (models.User, core.Error) {
	user := service.repository.FindByID(id)

	if user.ID == "" {
		return user, core.Error{"reason": "User not found"}
	}

	return user, nil
}
