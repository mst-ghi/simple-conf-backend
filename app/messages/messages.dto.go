package messages

type CreateDto struct {
	RoomID  string `json:"room_id" binding:"required"`
	Content string `json:"content" binding:"required,min=1,max=3000"`
}

type UpdateDto struct {
	Content string `json:"content" binding:"required,min=1,max=3000"`
}
