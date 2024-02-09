package users

import (
	"video-conf/core"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	root    *core.Controller
	service UsersServiceInterface
}

func NewUsersController() *UsersController {
	return &UsersController{
		root:    core.GetController(),
		service: NewUsersService(),
	}
}

// @tags     Users
// @security Bearer
// @router   /api/v1/users [get]
// @summary  get list of users
// @accept   json
// @produce  json
// @success  200 {object} core.Response[UsersResponseType]
func (ctrl *UsersController) FindAll(c *gin.Context) {
	user := core.User(c)
	users := ctrl.service.FindAll(user.ID)
	ctrl.root.Success(c, UsersResponse(users))
}

// @tags     Users
// @security Bearer
// @router   /api/v1/users/{id} [get]
// @summary  get user by id
// @accept   json
// @produce  json
// @success  200 {object} core.Response[UserResponseType]
// @param    id path string true "User ID"
func (ctrl *UsersController) FindOne(c *gin.Context) {
	user, err := ctrl.service.FindOne(c.Param("id"))

	if err != nil {
		ctrl.root.NotFoundError(c, err)
		return
	}

	ctrl.root.Success(c, UserResponse(user))
}
