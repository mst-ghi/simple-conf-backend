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
func (ctrl *CommunitiesController) FindAll(c *gin.Context) {
	communities := ctrl.service.FindAll()
	ctrl.root.Success(c, CommunitiesResponse(communities))
}

// @tags     Communities
// @security Bearer
// @router   /api/v1/communities/own [get]
// @summary  get list of own communities
// @accept   json
// @produce  json
// @success  200 {object} core.Response[CommunitiesResponseType]
func (ctrl *CommunitiesController) OwnerCommunities(c *gin.Context) {
	user := core.User(c)
	communities := ctrl.service.OwnerCommunities(user.ID)
	ctrl.root.Success(c, CommunitiesResponse(communities))
}

// @tags     Communities
// @security Bearer
// @router   /api/v1/communities/joined [get]
// @summary  get list of joined communities
// @accept   json
// @produce  json
// @success  200 {object} core.Response[CommunitiesResponseType]
func (ctrl *CommunitiesController) JoinedCommunities(c *gin.Context) {
	user := core.User(c)
	communities := ctrl.service.JoinedCommunities(user.ID)
	ctrl.root.Success(c, CommunitiesResponse(communities))
}

// @tags    Communities
// @router  /api/v1/communities/{id} [get]
// @summary get community by id
// @accept  json
// @produce json
// @success 200 {object} core.Response[CommunityResponseType]
// @param   id path string true "Community ID"
func (ctrl *CommunitiesController) FindOne(c *gin.Context) {
	community, err := ctrl.service.FindOne(c.Param("id"))

	if err != nil {
		ctrl.root.NotFoundError(c, err)
		return
	}

	ctrl.root.Success(c, CommunityResponse(community))
}

// @tags     Communities
// @security Bearer
// @router   /api/v1/communities [post]
// @summary  create new community
// @accept   json
// @produce  json
// @success  200 {object} core.Response[CommunityResponseType]
// @param    request body CreateDto true "Create community inputs"
func (ctrl *CommunitiesController) Create(c *gin.Context) {
	var dto CreateDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		ctrl.root.JsonBindError(c, err)
		return
	}

	user := core.User(c)
	community := ctrl.service.Create(user.ID, dto)

	ctrl.root.Success(c, CommunityResponse(community))
}

// @tags     Communities
// @security Bearer
// @router   /api/v1/communities/{id} [put]
// @summary  update community
// @accept   json
// @produce  json
// @success  200 {object} core.SuccessResponse
// @param    id path string true "Community ID"
// @param    request body UpdateDto true "Update community inputs"
func (ctrl *CommunitiesController) Update(c *gin.Context) {
	var dto UpdateDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		ctrl.root.JsonBindError(c, err)
		return
	}

	user := core.User(c)
	communityId := c.Param("id")

	err := ctrl.service.Update(user.ID, communityId, dto)

	if err != nil {
		ctrl.root.NotFoundError(c, err)
		return
	}

	ctrl.root.Success(c, nil)
}

// @tags     Communities
// @security Bearer
// @router   /api/v1/communities/{id}/join [put]
// @summary  join to community
// @accept   json
// @produce  json
// @success  200 {object} core.SuccessResponse
// @param    id path string true "Community ID"
func (ctrl *CommunitiesController) Join(c *gin.Context) {
	user := core.User(c)
	communityId := c.Param("id")

	err := ctrl.service.Join(communityId, user.ID)

	if err != nil {
		ctrl.root.BadRequestError(c, err)
		return
	}

	ctrl.root.Success(c, nil)
}

// @tags     Communities
// @security Bearer
// @router   /api/v1/communities/{id}/left [put]
// @summary  left from community
// @accept   json
// @produce  json
// @success  200 {object} core.SuccessResponse
// @param    id path string true "Community ID"
func (ctrl *CommunitiesController) Left(c *gin.Context) {
	user := core.User(c)
	communityId := c.Param("id")

	err := ctrl.service.Left(communityId, user.ID)

	if err != nil {
		ctrl.root.BadRequestError(c, err)
		return
	}

	ctrl.root.Success(c, nil)
}
