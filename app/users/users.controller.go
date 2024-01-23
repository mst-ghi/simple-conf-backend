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

// @tags    Users
// @router  /api/v1/users [get]
// @summary get list of users
// @accept  json
// @produce json
// @success 200 {object} core.Response[UsersResponseType]
func (self *UsersController) FindAll(c *gin.Context) {
	users := self.service.FindAll()
	self.root.Success(c, UsersResponse(users))
}

// @tags    Users
// @router  /api/v1/users/{id} [get]
// @summary get user by id
// @accept  json
// @produce json
// @success 200 {object} core.Response[UserResponseType]
// @param   id path int true "User ID"
func (self *UsersController) FindOne(c *gin.Context) {
	user, err := self.service.FindOne(c.Param("id"))

	if err != nil {
		self.root.NotFoundError(c, err)
		return
	}

	self.root.Success(c, UserResponse(user))
}
