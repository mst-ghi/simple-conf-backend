package models

import (
	"time"
	"video-conf/pkg/helpers"

	"gorm.io/gorm"
)

const (
	EVENT_PENDING_STATUS  = "pending"
	EVENT_STARTED_STATUS  = "started"
	EVENT_FINISHED_STATUS = "finished"
)

type Event struct {
	ID          string `gorm:"type:varchar(40);primarykey"`
	CommunityID string `gorm:"type:varchar(40);not null"`
	Title       string `gorm:"type:varchar(191);not null"`
	Description string `gorm:"type:varchar(255);not null"`
	Duration    uint8  `gorm:"type:SMALLINT UNSIGNED;size:240;not null"`
	Status      string `gorm:"type:varchar(40);default:pending"`

	Community Community `gorm:"foreignkey:CommunityID"`

	StartAt   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *Event) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = helpers.CUID()
	return
}
