package comments

import (
	"simple-conf/core"
	"simple-conf/database/models"
	"simple-conf/database/repositories"
)

type CommentsServiceInterface interface {
	FindAll(modelId, modelType string) ([]models.Comment, core.Error)
	Create(userId string, dto CreateDto) models.Comment
	Update(userId, id string, dto UpdateDto) core.Error
}

type CommentsService struct {
	repository repositories.CommentRepositoryInterface
}

func NewCommentsService() *CommentsService {
	return &CommentsService{
		repository: repositories.NewCommentRepository(),
	}
}

func (service *CommentsService) FindAll(modelId, modelType string) ([]models.Comment, core.Error) {
	if modelId == "" || modelType == "" {
		return []models.Comment{}, core.Error{"reason": "model id & model type is required"}
	}
	return service.repository.FindByModelID(modelId, modelType), nil
}

func (service *CommentsService) Create(userId string, dto CreateDto) models.Comment {
	comment := models.Comment{
		UserID:    userId,
		ModelID:   dto.ModelID,
		ModelType: dto.ModelType,
		Content:   dto.Content,
	}
	return service.repository.Create(comment)
}

func (service *CommentsService) Update(userId, id string, dto UpdateDto) core.Error {
	comment := service.repository.FindByID(id, userId)

	if comment.ID == "" {
		return core.Error{"reason": "Comment not found"}
	}

	service.repository.
		Connection().
		Table("comments").
		Where("id = ?", comment.ID).
		Updates(
			models.Comment{
				Content: dto.Content,
			},
		)

	return nil
}
