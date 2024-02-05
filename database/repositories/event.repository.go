package repositories

import (
	"video-conf/database"
	"video-conf/database/models"

	"gorm.io/gorm"
)

type EventRepositoryInterface interface {
	Connection() *gorm.DB
	Create(event models.Event) models.Event
	FindByID(id string) models.Event
	FindAll(eventId string) []models.Event
	Delete(eventId string)
}

type EventRepository struct {
	DB *gorm.DB
}

func NewEventRepository() *EventRepository {
	return &EventRepository{
		DB: database.Connection(),
	}
}

func (repo *EventRepository) Connection() *gorm.DB {
	return repo.DB
}

func (repo *EventRepository) Create(event models.Event) models.Event {
	var newEvent models.Event
	repo.DB.Create(&event).Scan(&newEvent)
	return newEvent
}

func (repo *EventRepository) FindByID(id string) models.Event {
	var event models.Event

	repo.DB.Preload("Community.Owner").First(&event, "id = ?", id)

	return event
}

func (repo *EventRepository) FindAll(communityId string) []models.Event {
	var events []models.Event

	query := repo.DB.Table("events")

	if communityId != "" {
		query = query.Where("community_id = ?", communityId)
	}

	query.Preload("Community.Owner").Order("created_at desc").Find(&events)

	return events
}

func (repo *EventRepository) Delete(eventId string) {
	repo.DB.
		Exec(
			"DELETE FROM events WHERE id = ?",
			eventId,
		)
}
