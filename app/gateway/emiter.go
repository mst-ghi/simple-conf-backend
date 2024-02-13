package gateway

import (
	"encoding/json"
)

func GetEmptyResponseData(code int, message string) string {
	data := SocketData{
		Code:    code,
		Message: message,
		Errors:  struct{}{},
		Data:    struct{}{},
	}

	result, _ := json.Marshal(data)

	return string(result)
}

func SuccessResponse(data interface{}) SocketData {
	return SocketData{
		Code:    SOCKET_STATUS_OK,
		Message: "Successful processing",
		Errors:  struct{}{},
		Data:    data,
	}
}

func ErrorResponse(code int, errors interface{}, data interface{}) SocketData {
	return SocketData{
		Code:    code,
		Message: "Successful processing",
		Errors:  errors,
		Data:    data,
	}
}

func BroadcastToGeneral(event string, data SocketData) {
	socket.BroadcastToRoom("", SOCKET_GENERAL_ROOM, event, data)
}
