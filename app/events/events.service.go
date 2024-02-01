package events

import (
	"video-conf/core"
	"video-conf/database/models"
	"video-conf/database/repositories"
)

type EventsServiceInterface interface {
	FindAll(communityID string) []models.Event
	FindOne(id string) (models.Event, core.Error)
	Create(ownerId string, dto CreateDto) models.Event
	Update(ownerId, id string, dto UpdateDto) core.Error
}

type EventsService struct {
	repository          repositories.EventRepositoryInterface
	communityRepository repositories.CommunityRepositoryInterface
}

func NewEventsService() *EventsService {
	return &EventsService{
		repository:          repositories.NewEventRepository(),
		communityRepository: repositories.NewCommunityRepository(),
	}
}

func (service *EventsService) FindAll(communityID string) []models.Event {
	return service.repository.FindAll(communityID)
}

func (service *EventsService) FindOne(id string) (models.Event, core.Error) {
	event := service.repository.FindByID(id)

	if event.ID == "" {
		return event, core.Error{"reason": "Event not found"}
	}

	return event, nil
}

func (service *EventsService) Create(ownerId string, dto CreateDto) models.Event {
	event := models.Event{
		CommunityID: dto.CommunityID,
		Title:       dto.Title,
		Description: dto.Description,
		Duration:    dto.Duration,
		Status:      dto.Status,
		StartAt:     dto.StartAt,
	}
	return service.repository.Create(event)
}

func (service *EventsService) Update(ownerId, id string, dto UpdateDto) core.Error {
	event := service.repository.FindByID(id)

	if event.ID == "" {
		return core.Error{"reason": "Event not found"}
	}

	community := service.communityRepository.FindByIDAndOwnerID(event.CommunityID, ownerId)

	if community.ID == "" {
		return core.Error{"reason": "You do not have access to edit this event"}
	}

	service.repository.
		Connection().
		Model(&event).
		Updates(
			models.Event{
				Title:       dto.Title,
				Description: dto.Description,
				Duration:    dto.Duration,
				Status:      dto.Status,
				StartAt:     dto.StartAt,
			},
		)

	return nil
}
