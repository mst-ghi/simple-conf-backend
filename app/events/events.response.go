package events

import (
	"simple-conf/app/communities"
	"simple-conf/database/models"
	"simple-conf/database/scopes"
	"time"
)

type ResponseType map[string]any

type Event struct {
	ID          string                `json:"id"`
	CommunityID string                `json:"community_id"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Duration    uint8                 `json:"duration"`
	Mode        string                `json:"mode"`
	Status      string                `json:"status"`
	StartAt     string                `json:"start_at"`
	CreatedAt   string                `json:"created_at"`
	UpdatedAt   string                `json:"updated_at"`
	Community   communities.Community `json:"community"`
}

func EventTransform(event models.Event) Event {
	if event.Community.ID != "" {
		return Event{
			ID:          event.ID,
			CommunityID: event.CommunityID,
			Title:       event.Title,
			Description: event.Description,
			Duration:    event.Duration,
			Mode:        event.Mode,
			Status:      event.Status,
			StartAt:     event.StartAt.Format(time.RFC3339),
			CreatedAt:   event.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   event.UpdatedAt.Format(time.RFC3339),
			Community:   communities.CommunityTransform(event.Community),
		}
	}
	return Event{
		ID:          event.ID,
		CommunityID: event.CommunityID,
		Title:       event.Title,
		Description: event.Description,
		Duration:    event.Duration,
		Mode:        event.Mode,
		Status:      event.Status,
		StartAt:     event.StartAt.Format(time.RFC3339),
		CreatedAt:   event.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   event.UpdatedAt.Format(time.RFC3339),
		Community:   communities.Community{},
	}
}

func EventsTransform(events []models.Event) []Event {
	data := []Event{}

	for _, event := range events {
		data = append(data, EventTransform(event))
	}

	return data
}

type EventResponseType struct {
	Event Event `json:"event"`
}

func EventResponse(event models.Event) ResponseType {
	return ResponseType{
		"event": EventTransform(event),
	}
}

type EventsResponseType struct {
	Events []Event `json:"events"`
}

func EventsResponse(events []models.Event) ResponseType {
	return ResponseType{
		"events": EventsTransform(events),
	}
}

type EventsMetaResponseType struct {
	Events []Event                 `json:"events"`
	Meta   scopes.PaginateMetadata `json:"meta"`
}

func EventsMetaResponse(events []models.Event, meta scopes.PaginateMetadata) ResponseType {
	return ResponseType{
		"events": EventsTransform(events),
		"meta":   meta,
	}
}
