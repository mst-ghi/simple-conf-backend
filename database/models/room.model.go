package models

import (
	"time"
	"video-conf/pkg/helpers"

	"gorm.io/gorm"
)

const (
	ROOM_MODE_PRIVATE  = "private"
	ROOM_MODE_PUBLIC   = "public"
	ROOM_ACCESS_OWNER  = "owner"
	ROOM_ACCESS_MEMBER = "member"
)

type Room struct {
	ID          string `gorm:"type:varchar(40);primarykey"`
	OwnerID     string `gorm:"type:varchar(40);not null"`
	Title       string `gorm:"type:varchar(191);not null"`
	Description string `gorm:"type:varchar(255)"`
	Mode        string `gorm:"type:varchar(40);default:private"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Owner User   `gorm:"foreignkey:OwnerID"`
	Users []User `gorm:"many2many:room_users"`
}

type RoomUser struct {
	RoomID    string `gorm:"primaryKey"`
	UserID    string `gorm:"primaryKey"`
	Access    string `gorm:"type:varchar(40);default:owner"`
	CreatedAt time.Time

	Room Room `gorm:"foreignkey:RoomID"`
	User User `gorm:"foreignkey:UserID"`
}

func (t *Room) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = helpers.CUID()
	return
}
