package comments

type CreateDto struct {
	ModelID   string `json:"model_id" binding:"required"`
	ModelType string `json:"model_type" binding:"required"`
	Content   string `json:"content" binding:"required,min=2,max=250"`
}

type UpdateDto struct {
	Content string `json:"content" binding:"required,min=2,max=250"`
}
