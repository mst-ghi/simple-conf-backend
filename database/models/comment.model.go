package models

import (
	"simple-conf/pkg/helpers"
	"time"

	"gorm.io/gorm"
)

const (
	COMMENT_TYPE_COMMUNITY = "community"
	COMMENT_TYPE_EVENT     = "event"
)

type Comment struct {
	ID        string `gorm:"type:varchar(40);primarykey"`
	UserID    string `gorm:"type:varchar(40);not null"`
	ModelID   string `gorm:"type:varchar(40);not null"`
	ModelType string `gorm:"type:varchar(40);not null"`
	Content   string `gorm:"type:varchar(3000);not null"`
	CreatedAt time.Time

	User User `gorm:"foreignkey:UserID"`
}

func (t *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = helpers.CUID()
	return
}
