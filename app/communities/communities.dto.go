package communities

type CreateDto struct {
	Title       string `json:"title" binding:"required,min=2,max=190"`
	Description string `json:"description" binding:"required,min=2,max=250"`
}

type UpdateDto struct {
	Title       string `json:"title" binding:"required,min=2,max=190"`
	Description string `json:"description" binding:"required,min=2,max=250"`
}
