package rooms

import (
	"simple-conf/core"
	"simple-conf/database/models"
	"simple-conf/database/repositories"
)

type RoomsServiceInterface interface {
	FindAll(userID string) []models.RoomUser
	FindOne(ownerId, id string) (models.RoomUser, core.Error)
	Create(ownerId string, dto CreateDto) (models.RoomUser, core.Error)
	Update(ownerId, id string, dto UpdateDto) core.Error
}

type RoomsService struct {
	repository          repositories.RoomRepositoryInterface
	communityRepository repositories.CommunityRepositoryInterface
}

func NewRoomsService() *RoomsService {
	return &RoomsService{
		repository:          repositories.NewRoomRepository(),
		communityRepository: repositories.NewCommunityRepository(),
	}
}

func (service *RoomsService) FindAll(userID string) []models.RoomUser {
	return service.repository.FindAll(userID)
}

func (service *RoomsService) FindOne(userId, id string) (models.RoomUser, core.Error) {
	room := service.repository.FindRoomUserByID(id, userId)

	if room.UserID == "" {
		return room, core.Error{"reason": "Room not found"}
	}

	return room, nil
}

func (service *RoomsService) Create(ownerId string, dto CreateDto) (models.RoomUser, core.Error) {
	room := models.Room{
		OwnerID:     ownerId,
		Title:       dto.Title,
		Description: dto.Description,
		Mode:        dto.Mode,
	}
	return service.repository.Create(room, dto.UserIDs)
}

func (service *RoomsService) Update(ownerId, id string, dto UpdateDto) core.Error {
	room := service.repository.FindRoomByID(id, ownerId)

	if room.ID == "" {
		return core.Error{"reason": "Room not found"}
	}

	service.repository.
		Connection().
		Table("rooms").
		Where("id = ?", room.ID).
		Updates(
			models.Room{
				Title:       dto.Title,
				Description: dto.Description,
				Mode:        dto.Mode,
			},
		)

	return nil
}
