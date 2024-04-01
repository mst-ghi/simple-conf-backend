package rooms

import (
	"simple-conf/app/users"
	"simple-conf/database/models"
	"time"
)

type ResponseType map[string]any

type RoomUser struct {
	ID          string `json:"id"`
	OwnerID     string `json:"owner_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Mode        string `json:"mode"`
	Access      string `json:"access"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`

	Users []users.UserShort `json:"users"`
}

func RoomUserTransform(room models.RoomUser) RoomUser {
	return RoomUser{
		ID:          room.RoomID,
		OwnerID:     room.Room.OwnerID,
		Title:       room.Room.Title,
		Description: room.Room.Description,
		Mode:        room.Room.Mode,
		Access:      room.Access,
		CreatedAt:   room.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   room.Room.UpdatedAt.Format(time.RFC3339),

		Users: users.UsersShortTransform(room.Room.Users),
	}
}

func RoomUsersTransform(rooms []models.RoomUser) []RoomUser {
	data := []RoomUser{}

	for _, room := range rooms {
		data = append(data, RoomUserTransform(room))
	}

	return data
}

type RoomResponseType struct {
	Room RoomUser `json:"room"`
}

func RoomResponse(room models.RoomUser) ResponseType {
	return ResponseType{
		"room": RoomUserTransform(room),
	}
}

type RoomsResponseType struct {
	Rooms []RoomUser `json:"rooms"`
}

func RoomsResponse(rooms []models.RoomUser) ResponseType {
	return ResponseType{
		"rooms": RoomUsersTransform(rooms),
	}
}
