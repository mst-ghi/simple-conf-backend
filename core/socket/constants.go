package socket

import "net/http"

const (
	SOCKET_STATUS_OK             = http.StatusOK
	SOCKET_STATUS_BAD_REQUEST    = http.StatusBadRequest
	SOCKET_STATUS_UNAUTHORIZED   = http.StatusUnauthorized
	SOCKET_STATUS_FORBIDDEN      = http.StatusForbidden
	SOCKET_STATUS_NOTFOUND       = http.StatusNotFound
	SOCKET_STATUS_UNPROCESSABLE  = http.StatusUnprocessableEntity
	SOCKET_STATUS_INTERNAL_ERROR = http.StatusInternalServerError
)

const (
	SOCKET_GENERAL_ROOM = "general"
)

const (
	EVENT_USER_ME  = "user:me"
	EVENT_USER_GET = "user:get"

	EVENT_ROOM_NEW    = "room:new"
	EVENT_ROOM_UPDATE = "room:update"
	EVENT_ROOM_DELETE = "room:delete"

	EVENT_MESSAGE_NEW    = "message:new"
	EVENT_MESSAGE_UPDATE = "message:update"
	EVENT_MESSAGE_DELETE = "message:delete"
)

type SocketUser struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type SocketContext struct {
	User SocketUser `json:"user"`
}

type SocketData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}
