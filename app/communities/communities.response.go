package communities

import (
	"time"
	"video-conf/app/users"
	"video-conf/database/models"
)

type ResponseType map[string]any

type Community struct {
	ID          string `json:"id"`
	OwnerID     string `json:"owner_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`

	Owner users.UserShort   `json:"owner"`
	Users []users.UserShort `json:"users"`
}

func CommunityTransform(community models.Community) Community {
	return Community{
		ID:          community.ID,
		OwnerID:     community.OwnerID,
		Title:       community.Title,
		Description: *community.Description,
		Status:      community.Status,
		CreatedAt:   community.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   community.UpdatedAt.Format(time.RFC3339),

		Owner: users.UserShortTransform(community.Owner),
		Users: users.UsersShortTransform(community.Users),
	}
}

func CommunitiesTransform(communities []models.Community) []Community {
	var data []Community

	for _, community := range communities {
		data = append(data, CommunityTransform(community))
	}

	return data
}

type CommunityResponseType struct {
	Community Community `json:"community"`
}

func CommunityResponse(community models.Community) ResponseType {
	return ResponseType{
		"community": CommunityTransform(community),
	}
}

type CommunitiesResponseType struct {
	Communities []Community `json:"communities"`
}

func CommunitiesResponse(communities []models.Community) ResponseType {
	return ResponseType{
		"communities": CommunitiesTransform(communities),
	}
}
