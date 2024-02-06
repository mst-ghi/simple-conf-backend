package rooms

type CreateDto struct {
	Title       string   `json:"title" binding:"required,min=2,max=190"`
	Description string   `json:"description" binding:"required,min=2,max=250"`
	Mode        string   `json:"mode" binding:"required"`
	UserIDs     []string `json:"user_ids" binding:"required,dive,min=1,max=100"`
}

type UpdateDto struct {
	Title       string `json:"title" binding:"required,min=2,max=190"`
	Description string `json:"description" binding:"required,min=2,max=250"`
	Mode        string `json:"mode" binding:"required"`
}
