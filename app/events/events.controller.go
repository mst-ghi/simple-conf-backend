package events

import (
	"fmt"
	"video-conf/core"

	"github.com/gin-gonic/gin"
)

type EventsController struct {
	root    *core.Controller
	service EventsServiceInterface
}

func NewEventsController() *EventsController {
	return &EventsController{
		root:    core.GetController(),
		service: NewEventsService(),
	}
}

// @tags    Events
// @router  /api/v1/events [get]
// @summary get list of events
// @accept  json
// @produce json
// @param   community_id query string false "Community ID"
// @success 200 {object} core.Response[EventsResponseType]
func (ctrl *EventsController) FindAll(c *gin.Context) {
	events := ctrl.service.FindAll(c.Query("community_id"))
	ctrl.root.Success(c, EventsResponse(events))
}

// @tags    Events
// @router  /api/v1/events/{id} [get]
// @summary get event by id
// @accept  json
// @produce json
// @success 200 {object} core.Response[EventResponseType]
// @param   id path string true "Event ID"
func (ctrl *EventsController) FindOne(c *gin.Context) {
	event, err := ctrl.service.FindOne(c.Param("id"))

	if err != nil {
		ctrl.root.NotFoundError(c, err)
		return
	}

	ctrl.root.Success(c, EventResponse(event))
}

// @tags     Events
// @security Bearer
// @router   /api/v1/events [post]
// @summary  create new event
// @accept   json
// @produce  json
// @success  200 {object} core.Response[EventResponseType]
// @param    request body CreateDto true "Create event inputs"
func (ctrl *EventsController) Create(c *gin.Context) {
	var dto CreateDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		fmt.Println(err)
		ctrl.root.JsonBindError(c, err)
		return
	}

	user := core.User(c)
	event := ctrl.service.Create(user.ID, dto)

	ctrl.root.Success(c, EventResponse(event))
}

// @tags     Events
// @security Bearer
// @router   /api/v1/events/{id} [put]
// @summary  update event
// @accept   json
// @produce  json
// @success  200 {object} core.SuccessResponse
// @param    id path string true "Event ID"
// @param    request body UpdateDto true "Update event inputs"
func (ctrl *EventsController) Update(c *gin.Context) {
	var dto UpdateDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		ctrl.root.JsonBindError(c, err)
		return
	}

	user := core.User(c)
	eventId := c.Param("id")

	err := ctrl.service.Update(user.ID, eventId, dto)

	if err != nil {
		ctrl.root.NotFoundError(c, err)
		return
	}

	ctrl.root.Success(c, nil)
}
