package communities

import (
	"video-conf/core"

	"github.com/gin-gonic/gin"
)

type CommunitiesController struct {
	root    *core.Controller
	service CommunitiesServiceInterface
}

func NewCommunitiesController() *CommunitiesController {
	return &CommunitiesController{
		root:    core.GetController(),
		service: NewCommunitiesService(),
	}
}

// @tags    Communities
// @router  /api/v1/communities [get]
// @summary get list of communities
// @accept  json
// @produce json
// @success 200 {object} core.Response[CommunitiesResponseType]
func (self *CommunitiesController) FindAll(c *gin.Context) {
	communities := self.service.FindAll()
	self.root.Success(c, CommunitiesResponse(communities))
}

// @tags    Communities
// @router  /api/v1/communities/{id} [get]
// @summary get community by id
// @accept  json
// @produce json
// @success 200 {object} core.Response[CommunityResponseType]
// @param   id path string true "Community ID"
func (self *CommunitiesController) FindOne(c *gin.Context) {
	community, err := self.service.FindOne(c.Param("id"))

	if err != nil {
		self.root.NotFoundError(c, err)
		return
	}

	self.root.Success(c, CommunityResponse(community))
}

// @tags     Communities
// @security Bearer
// @router   /api/v1/communities [post]
// @summary  create new community
// @accept   json
// @produce  json
// @success  200 {object} core.Response[CommunityResponseType]
// @param    request body CreateDto true "Create community inputs"
func (self *CommunitiesController) Create(c *gin.Context) {
	var dto CreateDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		self.root.JsonBindError(c, err)
		return
	}

	user := core.User(c)
	community := self.service.Create(user.ID, dto)

	self.root.Success(c, CommunityResponse(community))
}

// @tags     Communities
// @security Bearer
// @router   /api/v1/communities/{id} [put]
// @summary  update community
// @accept   json
// @produce  json
// @success  200 {object} core.SuccessResponse
// @param    id path string true "Community ID"
func (self *CommunitiesController) Update(c *gin.Context) {
	var dto UpdateDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		self.root.JsonBindError(c, err)
		return
	}

	user := core.User(c)
	communityId := c.Param("id")

	err := self.service.Update(user.ID, communityId, dto)

	if err != nil {
		self.root.NotFoundError(c, err)
		return
	}

	self.root.Success(c, nil)
}

// @tags     Communities
// @security Bearer
// @router   /api/v1/communities/{id}/join [put]
// @summary  join to community
// @accept   json
// @produce  json
// @success  200 {object} core.SuccessResponse
// @param    id path string true "Community ID"
func (self *CommunitiesController) Join(c *gin.Context) {
	user := core.User(c)
	communityId := c.Param("id")

	err := self.service.Join(communityId, user.ID)

	if err != nil {
		self.root.BadRequestError(c, err)
		return
	}

	self.root.Success(c, nil)
}
