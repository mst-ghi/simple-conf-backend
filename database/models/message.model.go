package models

import (
	"time"
	"video-conf/pkg/helpers"

	"gorm.io/gorm"
)

const (
	MESSAGE_TYPE_TEXT  = "text"
	MESSAGE_TYPE_LINK  = "link"
	MESSAGE_TYPE_IMAGE = "image"
	MESSAGE_TYPE_VIDEO = "video"
	MESSAGE_TYPE_File  = "file"
)

type Message struct {
	ID        string `gorm:"type:varchar(40);primarykey"`
	UserID    string `gorm:"type:varchar(40);not null"`
	RoomID    string `gorm:"type:varchar(40);not null"`
	Content   string `gorm:"type:varchar(3000);not null"`
	Type      string `gorm:"type:varchar(40);default:text"`
	CreatedAt time.Time
	UpdatedAt time.Time

	User User `gorm:"foreignkey:UserID"`
	Room Room `gorm:"foreignkey:RoomID"`
}

func (t *Message) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = helpers.CUID()
	return
}
