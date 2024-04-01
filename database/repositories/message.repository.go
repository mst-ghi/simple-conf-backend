package repositories

import (
	"simple-conf/database"
	"simple-conf/database/models"

	"gorm.io/gorm"
)

type MessageRepositoryInterface interface {
	Connection() *gorm.DB
	Create(message models.Message) models.Message
	FindByID(messageId, userID string) models.Message
	FindAll(roomId string) []models.Message
	Delete(messageId string)
}

type MessageRepository struct {
	DB *gorm.DB
}

func NewMessageRepository() *MessageRepository {
	return &MessageRepository{
		DB: database.Connection(),
	}
}

func (repo *MessageRepository) Connection() *gorm.DB {
	return repo.DB
}

func (repo *MessageRepository) Create(message models.Message) models.Message {
	var newMessage models.Message
	repo.DB.Create(&message).Preload("User").Scan(&newMessage)

	return repo.FindByID(newMessage.ID, newMessage.UserID)
}

func (repo *MessageRepository) FindByID(messageId, userID string) models.Message {
	var message models.Message

	repo.DB.Preload("User").
		Where("id = ?", messageId).
		Where("user_id = ?", userID).
		First(&message)

	return message
}

func (repo *MessageRepository) FindAll(roomId string) []models.Message {
	var messages []models.Message

	repo.DB.Table("messages").
		Where("room_id = ?", roomId).
		Preload("User").
		Order("created_at desc").
		Find(&messages)

	return messages
}

func (repo *MessageRepository) Delete(messageId string) {
	repo.DB.
		Exec(
			"DELETE FROM messages WHERE id = ?",
			messageId,
		)
}
