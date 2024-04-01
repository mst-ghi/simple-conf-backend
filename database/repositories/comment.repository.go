package repositories

import (
	"simple-conf/database"
	"simple-conf/database/models"

	"gorm.io/gorm"
)

type CommentRepositoryInterface interface {
	Connection() *gorm.DB
	FindByModelID(modelID, modelType string) []models.Comment
	FindByID(id, userId string) models.Comment
	Create(comment models.Comment) models.Comment
	Delete(commentId string)
}

type CommentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository() *CommentRepository {
	return &CommentRepository{
		DB: database.Connection(),
	}
}

func (repo *CommentRepository) Connection() *gorm.DB {
	return repo.DB
}

func (repo *CommentRepository) FindByModelID(modelID, modelType string) []models.Comment {
	var comments []models.Comment

	repo.DB.
		Table("comments").
		Preload("User").
		Where("model_id = ?", modelID).
		Where("model_type = ?", modelType).
		Find(&comments)

	return comments
}

func (repo *CommentRepository) FindByID(id, userId string) models.Comment {
	var comment models.Comment

	repo.DB.
		Table("comments").
		Preload("User").
		Where("id = ?", id).
		Where("user_id = ?", userId).
		First(&comment)

	return comment
}

func (repo *CommentRepository) Create(comment models.Comment) models.Comment {
	var newComment models.Comment
	repo.DB.Create(&comment).Scan(&newComment)
	return repo.FindByID(newComment.ID, newComment.UserID)
}

func (repo *CommentRepository) Delete(commentId string) {
	repo.DB.
		Exec(
			"DELETE FROM comments WHERE id = ?",
			commentId,
		)
}
