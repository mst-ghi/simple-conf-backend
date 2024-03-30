package comments

import (
	"time"
	"video-conf/app/users"
	"video-conf/database/models"
	"video-conf/database/scopes"
)

type ResponseType map[string]any

type Comment struct {
	ID        string     `json:"id"`
	UserID    string     `json:"user_id"`
	ModelID   string     `json:"model_id"`
	ModelType string     `json:"model_type"`
	Content   string     `json:"content"`
	CreatedAt string     `json:"created_at"`
	User      users.User `json:"user"`
}

func CommentTransform(comment models.Comment) Comment {
	return Comment{
		ID:        comment.ID,
		UserID:    comment.UserID,
		ModelID:   comment.ModelID,
		ModelType: comment.ModelType,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt.Format(time.RFC3339),
		User:      users.UserTransform(comment.User),
	}
}

func CommentsTransform(comments []models.Comment) []Comment {
	data := []Comment{}

	for _, comment := range comments {
		data = append(data, CommentTransform(comment))
	}

	return data
}

type CommentResponseType struct {
	Comment Comment `json:"comment"`
}

func CommentResponse(comment models.Comment) ResponseType {
	return ResponseType{
		"comment": CommentTransform(comment),
	}
}

type CommentsResponseType struct {
	Comments []Comment `json:"comments"`
}

func CommentsResponse(comments []models.Comment) ResponseType {
	return ResponseType{
		"comments": CommentsTransform(comments),
	}
}

type CommentsMetaResponseType struct {
	Comments []Comment               `json:"comments"`
	Meta     scopes.PaginateMetadata `json:"meta"`
}

func CommentsMetaResponse(comments []models.Comment, meta scopes.PaginateMetadata) ResponseType {
	return ResponseType{
		"comments": CommentsTransform(comments),
		"meta":     meta,
	}
}
