package models

import (
	"simple-conf/pkg/helpers"
	"time"

	"gorm.io/gorm"
)

const (
	COMMUNITY_ACTIVE_STATUS   = "active"
	COMMUNITY_INACTIVE_STATUS = "inactive"
)

type Community struct {
	ID          string  `gorm:"type:varchar(40);primarykey"`
	OwnerID     string  `gorm:"type:varchar(40);not null"`
	Title       string  `gorm:"type:varchar(191);not null"`
	Description *string `gorm:"type:varchar(255);default:null"`
	Status      string  `gorm:"type:varchar(40);default:active"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Owner User    `gorm:"foreignkey:OwnerID"`
	Users []User  `gorm:"many2many:community_users;"`
	Event []Event `gorm:"foreignKey:CommunityID"`
}

type CommunityUser struct {
	CommunityID string `json:"community_id"`
	UserID      string `json:"user_id"`
}

func (t *Community) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = helpers.CUID()
	return
}
