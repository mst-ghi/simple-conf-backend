package repositories

import (
	"simple-conf/database"
	"simple-conf/database/models"
	"simple-conf/database/scopes"

	"gorm.io/gorm"
)

type CommunityRepositoryInterface interface {
	Connection() *gorm.DB
	Create(community models.Community) models.Community
	FindByID(id string) models.Community
	FindByIDAndOwnerID(id, ownerId string) models.Community
	FindAll(search string, page, take int) ([]models.Community, scopes.PaginateMetadata)
	OwnerCommunities(ownerId string, search string) []models.Community
	JoinedCommunities(ownerId string) []models.Community
	FindOneCommunityUser(communityId, userId string) models.CommunityUser
	AppendCommunityUser(communityId, userId string) models.CommunityUser
	DeleteCommunityUser(communityId, userId string)
}

type CommunityRepository struct {
	DB *gorm.DB
}

func NewCommunityRepository() *CommunityRepository {
	return &CommunityRepository{
		DB: database.Connection(),
	}
}

func (repo *CommunityRepository) Connection() *gorm.DB {
	return repo.DB
}

func (repo *CommunityRepository) Create(community models.Community) models.Community {
	var newCommunity models.Community
	repo.DB.Create(&community).Scan(&newCommunity)
	return newCommunity
}

func (repo *CommunityRepository) FindByID(id string) models.Community {
	var community models.Community

	repo.DB.
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

func (repo *CommunityRepository) FindByIDAndOwnerID(id, ownerId string) models.Community {
	var community models.Community

	repo.DB.
		Where("id = ?", id).
		Where("owner_id = ?", ownerId).
		Preload("Owner", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id", "email", "name")
		}).
		Preload("Users", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id", "email", "name")
		}).
		First(&community)

	return community
}

func (repo *CommunityRepository) FindAll(search string, page, take int) ([]models.Community, scopes.PaginateMetadata) {
	var communities []models.Community

	query := repo.DB.
		Scopes(scopes.Paginate(page, take)).
		Table("communities")

	if search != "" {
		query.Where("title LIKE ?", "%"+search+"%")
	}

	query.
		Where("status = ?", models.COMMUNITY_ACTIVE_STATUS).
		Preload("Owner", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id", "email", "name")
		}).
		Preload("Users", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id", "email", "name")
		}).
		Order("created_at desc").
		Find(&communities)

	var totalRows int64
	repo.DB.Table("communities").
		Where("status = ?", models.COMMUNITY_ACTIVE_STATUS).
		Count(&totalRows)

	return communities, scopes.PaginateMeta(totalRows, page, take)
}

func (repo *CommunityRepository) OwnerCommunities(ownerId string, search string) []models.Community {
	var communities []models.Community

	query := repo.DB.
		Table("communities")

	if search != "" {
		query.Where("title LIKE ?", "%"+search+"%")
	}

	query.
		Where("owner_id = ?", ownerId).
		Preload("Owner", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id", "email", "name")
		}).
		Preload("Users", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id", "email", "name")
		}).
		Order("created_at desc").
		Find(&communities)

	return communities
}

func (repo *CommunityRepository) JoinedCommunities(ownerId string) []models.Community {
	var user models.User

	repo.DB.Table("users").
		Where("id = ?", ownerId).
		Preload("Communities", func(tx *gorm.DB) *gorm.DB {
			return tx.
				Where("status = ?", models.COMMUNITY_ACTIVE_STATUS).
				Preload("Owner", func(tx *gorm.DB) *gorm.DB {
					return tx.Select("id", "email", "name")
				}).
				Preload("Users", func(tx *gorm.DB) *gorm.DB {
					return tx.Select("id", "email", "name")
				})
		}).
		First(&user)

	return user.Communities
}

func (repo *CommunityRepository) FindOneCommunityUser(communityId, userId string) models.CommunityUser {
	var communityUser models.CommunityUser

	repo.DB.
		Raw(
			"SELECT * FROM community_users WHERE community_id = ? AND user_id = ?",
			communityId,
			userId,
		).
		Scan(&communityUser)

	return communityUser
}

func (repo *CommunityRepository) AppendCommunityUser(communityId, userId string) models.CommunityUser {
	var communityUser models.CommunityUser

	repo.DB.
		Raw(
			"INSERT INTO community_users (community_id, user_id) VALUES(?, ?)",
			communityId,
			userId,
		).
		Scan(&communityUser)

	return communityUser
}

func (repo *CommunityRepository) DeleteCommunityUser(communityId, userId string) {
	repo.DB.
		Exec(
			"DELETE FROM community_users WHERE community_id = ? AND user_id = ?",
			communityId,
			userId,
		)
}
