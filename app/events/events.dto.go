package events

import "time"

type CreateDto struct {
	CommunityID string    `json:"community_id" binding:"required"`
	Title       string    `json:"title" binding:"required,min=2,max=190"`
	Description string    `json:"description" binding:"required,min=2,max=250"`
	Duration    uint8     `json:"duration" binding:"required,gte=0,lte=250"`
	Status      string    `json:"status" binding:"required"`
	StartAt     time.Time `json:"start_at" binding:"required"`
}

type UpdateDto struct {
	Title       string    `json:"title" binding:"required,min=2,max=190"`
	Description string    `json:"description" binding:"required,min=2,max=250"`
	Duration    uint8     `json:"duration" binding:"required,gte=0,lte=250"`
	Status      string    `json:"status" binding:"required"`
	StartAt     time.Time `json:"start_at" binding:"required"`
}
