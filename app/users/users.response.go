package users

import (
	"time"
	"video-conf/database/models"
)

type ResponseType map[string]any

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func UserTransform(user models.User) User {
	return User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}
}

type UserResponseType struct {
	User User `json:"user"`
}

func UserResponse(user models.User) ResponseType {
	return ResponseType{
		"user": UserTransform(user),
	}
}

type UsersResponseType struct {
	Users []User `json:"users"`
}

func UsersResponse(users []models.User) ResponseType {
	var data []User

	for _, user := range users {
		data = append(data, UserTransform(user))
	}

	return ResponseType{
		"users": data,
	}
}
