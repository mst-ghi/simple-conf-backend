package communities

import (
	"video-conf/core"
	"video-conf/database/models"
	"video-conf/database/repositories"
)

type CommunitiesServiceInterface interface {
	FindAll() []models.Community
	FindOne(id string) (models.Community, core.Error)
	Create(ownerId string, dto CreateDto) models.Community
	Update(ownerId, id string, dto UpdateDto) core.Error
	Join(communityId, userId string) core.Error
}

type CommunitiesService struct {
	repository repositories.CommunityRepositoryInterface
}

func NewCommunitiesService() *CommunitiesService {
	return &CommunitiesService{
		repository: repositories.NewCommunityRepository(),
	}
}

func (self *CommunitiesService) FindAll() []models.Community {
	return self.repository.FindAll()
}

func (self *CommunitiesService) FindOne(id string) (models.Community, core.Error) {
	community := self.repository.FindByID(id)

	if community.ID == "" {
		return community, core.Error{"reason": "Community not found"}
	}

	return community, nil
}

func (self *CommunitiesService) Create(ownerId string, dto CreateDto) models.Community {
	community := models.Community{
		OwnerID:     ownerId,
		Title:       dto.Title,
		Description: &dto.Description,
		Status:      models.COMMUNITY_ACTIVE_STATUS,
	}
	return self.repository.Create(community)
}

func (self *CommunitiesService) Update(ownerId, id string, dto UpdateDto) core.Error {
	community := self.repository.FindByIDAndOwnerID(id, ownerId)

	if community.ID != "" {
		return core.Error{"reason": "Community not found"}
	}

	community.Title = dto.Title
	community.Description = &dto.Description

	self.repository.Connection().Save(community)

	return nil
}

func (self *CommunitiesService) Join(communityId, userId string) core.Error {
	community := self.repository.FindOneCommunityUser(communityId, userId)

	if community.CommunityID == communityId {
		return core.Error{"reason": "You join to community before"}
	}

	self.repository.AppendCommunityUser(communityId, userId)

	return nil
}
