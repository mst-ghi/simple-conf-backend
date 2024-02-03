package communities

import (
	"video-conf/core"
	"video-conf/database/models"
	"video-conf/database/repositories"
)

type CommunitiesServiceInterface interface {
	FindAll() []models.Community
	OwnerCommunities(ownerId string) []models.Community
	JoinedCommunities(ownerId string) []models.Community
	FindOne(id string) (models.Community, core.Error)
	Create(ownerId string, dto CreateDto) models.Community
	Update(ownerId, id string, dto UpdateDto) core.Error
	Join(communityId, userId string) core.Error
	Left(communityId, userId string) core.Error
}

type CommunitiesService struct {
	repository repositories.CommunityRepositoryInterface
}

func NewCommunitiesService() *CommunitiesService {
	return &CommunitiesService{
		repository: repositories.NewCommunityRepository(),
	}
}

func (service *CommunitiesService) FindAll() []models.Community {
	return service.repository.FindAll()
}

func (service *CommunitiesService) OwnerCommunities(ownerId string) []models.Community {
	return service.repository.OwnerCommunities(ownerId)
}

func (service *CommunitiesService) JoinedCommunities(ownerId string) []models.Community {
	return service.repository.JoinedCommunities(ownerId)
}

func (service *CommunitiesService) FindOne(id string) (models.Community, core.Error) {
	community := service.repository.FindByID(id)

	if community.ID == "" {
		return community, core.Error{"reason": "Community not found"}
	}

	return community, nil
}

func (service *CommunitiesService) Create(ownerId string, dto CreateDto) models.Community {
	community := models.Community{
		OwnerID:     ownerId,
		Title:       dto.Title,
		Description: &dto.Description,
		Status:      models.COMMUNITY_ACTIVE_STATUS,
	}
	return service.repository.Create(community)
}

func (service *CommunitiesService) Update(ownerId, id string, dto UpdateDto) core.Error {
	community := service.repository.FindByIDAndOwnerID(id, ownerId)

	if community.ID == "" {
		return core.Error{"reason": "Community not found"}
	}

	service.repository.
		Connection().
		Table("communities").
		Where("id = ?", community.ID).
		Updates(models.Community{Title: dto.Title, Description: &dto.Description})

	return nil
}

func (service *CommunitiesService) Join(communityId, userId string) core.Error {
	community := service.repository.FindOneCommunityUser(communityId, userId)

	if community.CommunityID == communityId {
		return core.Error{"reason": "You join to community before"}
	}

	service.repository.AppendCommunityUser(communityId, userId)

	return nil
}

func (service *CommunitiesService) Left(communityId, userId string) core.Error {
	community := service.repository.FindOneCommunityUser(communityId, userId)

	if community.CommunityID != "" {
		return core.Error{"reason": "You are not a member of the community"}
	}

	service.repository.DeleteCommunityUser(communityId, userId)

	return nil
}
