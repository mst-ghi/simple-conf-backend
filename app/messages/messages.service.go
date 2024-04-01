package messages

import (
	"simple-conf/core"
	"simple-conf/database/models"
	"simple-conf/database/repositories"
)

type MessagesServiceInterface interface {
	FindAll(roomID string) []models.Message
	Create(ownerId string, dto CreateDto) (models.Message, core.Error)
	Update(ownerId, id string, dto UpdateDto) core.Error
}

type MessagesService struct {
	repository     repositories.MessageRepositoryInterface
	roomRepository repositories.RoomRepositoryInterface
}

func NewMessagesService() *MessagesService {
	return &MessagesService{
		repository:     repositories.NewMessageRepository(),
		roomRepository: repositories.NewRoomRepository(),
	}
}

func (service *MessagesService) FindAll(roomID string) []models.Message {
	return service.repository.FindAll(roomID)
}

func (service *MessagesService) Create(ownerId string, dto CreateDto) (models.Message, core.Error) {
	var message models.Message
	var roomUser models.RoomUser

	service.repository.
		Connection().
		Table("room_users").
		Where("room_id", dto.RoomID).
		Where("user_id", ownerId).
		First(&roomUser)

	if roomUser.RoomID == "" {
		return message, core.Error{"reason": "You do not have access to do this"}
	}

	message = models.Message{
		RoomID:  dto.RoomID,
		UserID:  ownerId,
		Content: dto.Content,
		Type:    models.MESSAGE_TYPE_TEXT,
	}

	return service.repository.Create(message), nil
}

func (service *MessagesService) Update(ownerId, id string, dto UpdateDto) core.Error {
	message := service.repository.FindByID(id, ownerId)

	if message.ID == "" {
		return core.Error{"reason": "Message not found"}
	}

	service.repository.
		Connection().
		Table("messages").
		Where("id = ?", message.ID).
		Updates(
			models.Message{
				Content: dto.Content,
			},
		)

	return nil
}
