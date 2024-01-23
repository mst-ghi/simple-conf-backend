package auth

import (
	"video-conf/app/users"
	core "video-conf/core"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	root    *core.Controller
	service AuthServiceInterface
}

func NewAuthController() *AuthController {
	return &AuthController{
		root:    core.GetController(),
		service: NewAuthService(),
	}
}

// @tags    Auth
// @router  /api/v1/auth/login [post]
// @summary login api
// @accept  json
// @produce json
// @success 200 {object} core.Response[TokensResponseType]
// @param   request body LoginDto true "Login inputs"
func (self *AuthController) Login(c *gin.Context) {
	var data LoginDto

	if err := c.ShouldBindJSON(&data); err != nil {
		self.root.JsonBindError(c, err)
		return
	}

	tokens, err := self.service.Login(data)

	if err != nil {
		self.root.UnprocessableError(c, err)
		return
	}

	self.root.Success(c, TokensResponse(tokens))
}

// @tags    Auth
// @router  /api/v1/auth/register [post]
// @summary register api
// @accept  json
// @produce json
// @success 200 {object} core.SuccessResponse
// @param   request body RegisterDto true "Register inputs"
func (self *AuthController) Register(c *gin.Context) {
	var dto RegisterDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		self.root.JsonBindError(c, err)
		return
	}

	if err := self.service.Register(dto); err != nil {
		self.root.UnprocessableError(c, err)
		return
	}

	self.root.Success(c, nil)
}

// @tags    Auth
// @router  /api/v1/auth/refresh [post]
// @summary refresh tokens api
// @accept  json
// @produce json
// @success 200 {object} core.Response[TokensResponseType]
// @param   request body RefreshDto true "Refresh tokens inputs"
func (self *AuthController) Refresh(c *gin.Context) {
	var data RefreshDto

	if err := c.ShouldBindJSON(&data); err != nil {
		self.root.JsonBindError(c, err)
		return
	}

	tokens, err := self.service.Refresh(data)

	if err != nil {
		self.root.UnprocessableError(c, err)
		return
	}

	self.root.Success(c, TokensResponse(tokens))
}

// @tags     Auth
// @security Bearer
// @router   /api/v1/auth/user [get]
// @summary  fetch logged in user info
// @accept   json
// @produce  json
// @success  200 {object} core.Response[users.UserResponseType]
func (self *AuthController) User(c *gin.Context) {
	user := core.User(c)
	self.root.Success(c, users.UserResponse(user))
}

// @tags     Auth
// @security Bearer
// @router   /api/v1/auth/logout [get]
// @summary  logout user
// @accept   json
// @produce  json
// @success  200 {object} core.SuccessResponse
func (self *AuthController) Logout(c *gin.Context) {
	self.service.Logout(core.Token(c))
	self.root.Success(c, nil)
}
