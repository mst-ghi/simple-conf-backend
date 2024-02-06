package rooms

import (
	"video-conf/core"

	"github.com/gin-gonic/gin"
)

type RoomsController struct {
	root    *core.Controller
	service RoomsServiceInterface
}

func NewRoomsController() *RoomsController {
	return &RoomsController{
		root:    core.GetController(),
		service: NewRoomsService(),
	}
}

// @tags    Rooms
// @security Bearer
// @router  /api/v1/rooms [get]
// @summary get list of rooms
// @accept  json
// @produce json
// @success 200 {object} core.Response[RoomsResponseType]
func (ctrl *RoomsController) FindAll(c *gin.Context) {
	user := core.User(c)
	rooms := ctrl.service.FindAll(user.ID)

	ctrl.root.Success(c, RoomsResponse(rooms))
}

// @tags    Rooms
// @security Bearer
// @router  /api/v1/rooms/{id} [get]
// @summary get room by id
// @accept  json
// @produce json
// @success 200 {object} core.Response[RoomResponseType]
// @param   id path string true "Room ID"
func (ctrl *RoomsController) FindOne(c *gin.Context) {
	user := core.User(c)
	room, err := ctrl.service.FindOne(user.ID, c.Param("id"))

	if err != nil {
		ctrl.root.NotFoundError(c, err)
		return
	}

	ctrl.root.Success(c, RoomResponse(room))
}

// @tags     Rooms
// @security Bearer
// @router   /api/v1/rooms [post]
// @summary  create new room
// @accept   json
// @produce  json
// @success  200 {object} core.Response[RoomResponseType]
// @param    request body CreateDto true "Create room inputs"
func (ctrl *RoomsController) Create(c *gin.Context) {
	var dto CreateDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		ctrl.root.JsonBindError(c, err)
		return
	}

	user := core.User(c)
	roomUser, err := ctrl.service.Create(user.ID, dto)

	if err != nil {
		ctrl.root.BadRequestError(c, err)
		return
	}

	ctrl.root.Success(c, RoomResponse(roomUser))
}

// @tags     Rooms
// @security Bearer
// @router   /api/v1/rooms/{id} [put]
// @summary  update room
// @accept   json
// @produce  json
// @success  200 {object} core.SuccessResponse
// @param    id path string true "Room ID"
// @param    request body UpdateDto true "Update room inputs"
func (ctrl *RoomsController) Update(c *gin.Context) {
	var dto UpdateDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		ctrl.root.JsonBindError(c, err)
		return
	}

	user := core.User(c)
	roomId := c.Param("id")

	err := ctrl.service.Update(user.ID, roomId, dto)

	if err != nil {
		ctrl.root.NotFoundError(c, err)
		return
	}

	ctrl.root.Success(c, nil)
}
