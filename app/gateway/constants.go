package gateway

const (
	SOCKET_GENERAL_ROOM = "general"
)

const (
	EVENT_ERROR_UNAUTHORIZED = "errors:unauthorized"
	EVENT_ERROR_MESSAGE      = "errors:message"

	EVENT_USER_ME  = "user:me"
	EVENT_USER_GET = "user:get"

	EVENT_ROOM_NEW    = "room:new"
	EVENT_ROOM_UPDATE = "room:update"
	EVENT_ROOM_DELETE = "room:delete"

	EVENT_MESSAGE_SEND   = "message:send"
	EVENT_MESSAGE_NEW    = "message:new"
	EVENT_MESSAGE_UPDATE = "message:update"
	EVENT_MESSAGE_DELETE = "message:delete"

	EVENT_CALL_CALLING   = "call:calling"
	EVENT_CALL_RECEIVING = "call:receiving"
	EVENT_CALL_ACCEPTING = "call:accepting"
	EVENT_CALL_ACCEPTED  = "call:accepted"
	EVENT_CALL_ENDING    = "call:ending"
	EVENT_CALL_ENDED     = "call:ended"
	EVENT_CALL_BUSY      = "call:busy"
)
