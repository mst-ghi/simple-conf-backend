package gateway

import (
	"video-conf/app/messages"

	socketio "github.com/googollee/go-socket.io"
)

func ModuleHandlers() {
	messageGateway := messages.NewMessagesGateway()

	socket.OnEvent("/", EVENT_USER_GET, func(con socketio.Conn) {
		ctx, ok := CheckContext(con)

		if ok {
			SuccessEmitTo(con, EVENT_USER_ME, ctx)
		}

	})

	socket.OnEvent("/", EVENT_MESSAGE_SEND, func(con socketio.Conn, data map[string]string) {
		ctx, ok := CheckContext(con)

		if ok {
			message, err := messageGateway.NewMessage(
				messages.NewMessageData{
					UserId:  ctx.User.ID,
					RoomId:  data["RoomId"],
					Content: data["Content"],
				},
			)

			if err != nil {
				ErrorEmitTo(con, EVENT_ERROR_MESSAGE, SOCKET_STATUS_BAD_REQUEST, struct{}{})
			} else {
				BroadcastToRoom(
					message.RoomID,
					EVENT_MESSAGE_NEW,
					map[string]any{"message": messages.MessageTransform(message)},
				)
			}
		}
	})
}
