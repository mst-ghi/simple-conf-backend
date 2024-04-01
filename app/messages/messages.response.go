package messages

import (
	"simple-conf/app/users"
	"simple-conf/database/models"
	"time"
)

type ResponseType map[string]any

type Message struct {
	ID        string          `json:"id"`
	UserID    string          `json:"user_id"`
	RoomID    string          `json:"room_id"`
	Content   string          `json:"content"`
	Type      string          `json:"type"`
	CreatedAt string          `json:"created_at"`
	UpdatedAt string          `json:"updated_at"`
	User      users.UserShort `json:"user"`
}

func MessageTransform(message models.Message) Message {
	return Message{
		ID:        message.ID,
		UserID:    message.UserID,
		RoomID:    message.RoomID,
		Content:   message.Content,
		Type:      message.Type,
		CreatedAt: message.CreatedAt.Format(time.RFC3339),
		UpdatedAt: message.UpdatedAt.Format(time.RFC3339),
		User:      users.UserShortTransform(message.User),
	}
}

func MessagesTransform(messages []models.Message) []Message {
	data := []Message{}

	for _, message := range messages {
		data = append(data, MessageTransform(message))
	}

	return data
}

type MessageResponseType struct {
	Message Message `json:"message"`
}

func MessageResponse(message models.Message) ResponseType {
	return ResponseType{
		"message": MessageTransform(message),
	}
}

type MessagesResponseType struct {
	Messages []Message `json:"messages"`
}

func MessagesResponse(messages []models.Message) ResponseType {
	return ResponseType{
		"messages": MessagesTransform(messages),
	}
}
