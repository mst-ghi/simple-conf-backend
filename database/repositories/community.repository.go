package repositories

import (
	"video-conf/database"
	"video-conf/database/models"

	"gorm.io/gorm"
)

type CommunityRepositoryInterface interface {
	Connection() *gorm.DB
	Create(community models.Community) models.Community
	FindByID(id string) models.Community
	FindByIDAndOwnerID(id, ownerId string) models.Community
	FindAll() []models.Community
	FindOneCommunityUser(communityId, userId string) models.CommunityUser
	AppendCommunityUser(communityId, userId string) models.CommunityUser
}

type CommunityRepository struct {
	DB *gorm.DB
}

func NewCommunityRepository() *CommunityRepository {
	return &CommunityRepository{
		DB: database.Connection(),
	}
}

func (self *CommunityRepository) Connection() *gorm.DB {
	return self.DB
}

func (self *CommunityRepository) Create(community models.Community) models.Community {
	var newCommunity models.Community
	self.DB.Create(&community).Scan(&newCommunity)
	return newCommunity
}

func (self *CommunityRepository) FindByID(id string) models.Community {
	var community models.Community

	self.DB.
		Where("status = ?", models.COMMUNITY_ACTIVE_STATUS).
		Preload("Owner", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id", "email", "name")
		}).
		Preload("Users", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id", "email", "name")
		}).
		First(&community, "id = ?", id)

	return community
}

func (self *CommunityRepository) FindByIDAndOwnerID(id, ownerId string) models.Community {
	var community models.Community

	self.DB.
		Where("id = ?", id).Where("owner_id = ?", ownerId).
		Preload("Owner", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id", "email", "name")
		}).
		Preload("Users", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id", "email", "name")
		}).
		First(&community)

	return community
}

func (self *CommunityRepository) FindAll() []models.Community {
	var communities []models.Community

	self.DB.
		Table("communities").
		Where("status = ?", models.COMMUNITY_ACTIVE_STATUS).
		Preload("Owner", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id", "email", "name")
		}).
		Preload("Users", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id", "email", "name")
		}).
		Find(&communities)

	return communities
}

func (self *CommunityRepository) FindOneCommunityUser(communityId, userId string) models.CommunityUser {
	var communityUser models.CommunityUser

	self.DB.
		Raw(
			"SELECT * FROM community_users WHERE community_id = ? AND user_id = ?",
			communityId,
			userId,
		).
		Scan(&communityUser)

	return communityUser
}

func (self *CommunityRepository) AppendCommunityUser(communityId, userId string) models.CommunityUser {
	var communityUser models.CommunityUser

	self.DB.
		Raw(
			"INSERT INTO community_users (community_id, user_id) VALUES(?, ?)",
			communityId,
			userId,
		).
		Scan(&communityUser)

	return communityUser
}
