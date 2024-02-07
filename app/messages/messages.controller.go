package messages

import (
	"video-conf/core"

	"github.com/gin-gonic/gin"
)

type MessagesController struct {
	root    *core.Controller
	service MessagesServiceInterface
}

func NewMessagesController() *MessagesController {
	return &MessagesController{
		root:    core.GetController(),
		service: NewMessagesService(),
	}
}

// @tags    Messages
// @security Bearer
// @router  /api/v1/messages [get]
// @summary get list of messages
// @accept  json
// @produce json
// @param   room_id query string true "Room ID"
// @success 200 {object} core.Response[MessagesResponseType]
func (ctrl *MessagesController) FindAll(c *gin.Context) {
	messages := ctrl.service.FindAll(c.Query("room_id"))
	ctrl.root.Success(c, MessagesResponse(messages))
}

// @tags     Messages
// @security Bearer
// @router   /api/v1/messages [post]
// @summary  create new message
// @accept   json
// @produce  json
// @success  200 {object} core.Response[MessageResponseType]
// @param    request body CreateDto true "Create message inputs"
func (ctrl *MessagesController) Create(c *gin.Context) {
	var dto CreateDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		ctrl.root.JsonBindError(c, err)
		return
	}

	user := core.User(c)
	message, err := ctrl.service.Create(user.ID, dto)

	if err != nil {
		ctrl.root.BadRequestError(c, err)
		return
	}

	ctrl.root.Success(c, MessageResponse(message))
}

// @tags     Messages
// @security Bearer
// @router   /api/v1/messages/{id} [put]
// @summary  update message
// @accept   json
// @produce  json
// @success  200 {object} core.SuccessResponse
// @param    id path string true "Message ID"
// @param    request body UpdateDto true "Update message inputs"
func (ctrl *MessagesController) Update(c *gin.Context) {
	var dto UpdateDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		ctrl.root.JsonBindError(c, err)
		return
	}

	user := core.User(c)
	messageId := c.Param("id")

	err := ctrl.service.Update(user.ID, messageId, dto)

	if err != nil {
		ctrl.root.NotFoundError(c, err)
		return
	}

	ctrl.root.Success(c, nil)
}
