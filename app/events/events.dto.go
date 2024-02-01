package events

import "time"

type CreateDto struct {
	CommunityID string    `json:"community_id" binding:"required"`
	Title       string    `json:"title" binding:"required,min=2,max=190"`
	Description string    `json:"description" binding:"required,min=2,max=250"`
	Duration    uint8     `json:"duration" binding:"required,gte=1,lte=250"`
	Status      string    `json:"status" binding:"required,enum" enum:"pending,started,finished"`
	StartAt     time.Time `json:"start_at" binding:"required,datetime"`
}

type UpdateDto struct {
	Title       string    `json:"title" binding:"required,min=2,max=190"`
	Description string    `json:"description" binding:"required,min=2,max=250"`
	Duration    uint8     `json:"duration" binding:"required,gte=1,lte=250"`
	Status      string    `json:"status" binding:"required,enum" enum:"pending,started,finished"`
	StartAt     time.Time `json:"start_at" binding:"required,datetime"`
}
