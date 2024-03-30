package comments

import (
	"video-conf/core"

	"github.com/gin-gonic/gin"
)

type CommentsController struct {
	root    *core.Controller
	service CommentsServiceInterface
}

func NewCommentsController() *CommentsController {
	return &CommentsController{
		root:    core.GetController(),
		service: NewCommentsService(),
	}
}

// @tags    Comments
// @router  /api/v1/comments [get]
// @summary get list of comments
// @accept  json
// @produce json
// @param   model_id query string true "Model ID"
// @param   model_type query string true "Model type"
// @success 200 {object} core.Response[CommentsResponseType]
func (ctrl *CommentsController) FindAll(c *gin.Context) {
	comments, err := ctrl.service.FindAll(c.Query("model_id"), c.Query("model_type"))

	if err != nil {
		ctrl.root.BadRequestError(c, err)
		return
	}

	ctrl.root.Success(c, CommentsResponse(comments))
}

// @tags     Comments
// @security Bearer
// @router   /api/v1/comments [post]
// @summary  create new comment
// @accept   json
// @produce  json
// @success  200 {object} core.Response[CommentResponseType]
// @param    request body CreateDto true "Create comment inputs"
func (ctrl *CommentsController) Create(c *gin.Context) {
	var dto CreateDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		ctrl.root.JsonBindError(c, err)
		return
	}

	user := core.User(c)
	comment := ctrl.service.Create(user.ID, dto)

	ctrl.root.Success(c, CommentResponse(comment))
}

// @tags     Comments
// @security Bearer
// @router   /api/v1/comments/{id} [put]
// @summary  update comment
// @accept   json
// @produce  json
// @success  200 {object} core.SuccessResponse
// @param    id path string true "Comment ID"
// @param    request body UpdateDto true "Update comment inputs"
func (ctrl *CommentsController) Update(c *gin.Context) {
	var dto UpdateDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		ctrl.root.JsonBindError(c, err)
		return
	}

	user := core.User(c)
	commentId := c.Param("id")

	err := ctrl.service.Update(user.ID, commentId, dto)

	if err != nil {
		ctrl.root.NotFoundError(c, err)
		return
	}

	ctrl.root.Success(c, nil)
}
