package messages

import (
	"video-conf/core"
	"video-conf/database/models"
)

type MessagesGateway struct {
	service MessagesServiceInterface
}

func NewMessagesGateway() *MessagesGateway {
	return &MessagesGateway{
		service: NewMessagesService(),
	}
}

type NewMessageData struct{ UserId, RoomId, Content string }

func (gate *MessagesGateway) NewMessage(data NewMessageData) (models.Message, core.Error) {
	return gate.service.Create(
		data.UserId,
		CreateDto{RoomID: data.RoomId, Content: data.Content},
	)
}
