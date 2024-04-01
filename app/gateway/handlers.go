package gateway

import (
	"simple-conf/app/messages"

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
				ErrorEmitTo(con, EVENT_ERROR_MESSAGE, struct{}{})
			} else {
				BroadcastToRoom(
					message.RoomID,
					EVENT_MESSAGE_NEW,
					map[string]any{"message": messages.MessageTransform(message)},
				)
			}
		}
	})

	socket.OnEvent("/", EVENT_CALL_CALLING, func(con socketio.Conn, data map[string]any) {
		_, ok := CheckContext(con)

		if ok {
			userStatus := socket.RoomLen("/", data["toRoomId"].(string))

			if userStatus > 0 {
				BroadcastToRoom(
					data["toRoomId"].(string),
					EVENT_CALL_RECEIVING,
					data,
				)
			} else {
				BroadcastToRoom(
					data["fromRoomId"].(string),
					EVENT_CALL_OFFLINE,
					data,
				)
			}

		}
	})

	socket.OnEvent("/", EVENT_CALL_BUSY, func(con socketio.Conn, data map[string]any) {
		_, ok := CheckContext(con)

		if ok {
			BroadcastToRoom(
				data["fromRoomId"].(string),
				EVENT_CALL_BUSY,
				data,
			)
		}
	})

	socket.OnEvent("/", EVENT_CALL_ACCEPTING, func(con socketio.Conn, data map[string]any) {
		_, ok := CheckContext(con)

		if ok {
			BroadcastToRoom(
				data["fromRoomId"].(string),
				EVENT_CALL_ACCEPTED,
				data,
			)
		}
	})

	socket.OnEvent("/", EVENT_CALL_ENDING, func(con socketio.Conn, data map[string]any) {
		_, ok := CheckContext(con)

		if ok {
			BroadcastToRoom(
				data["toRoomId"].(string),
				EVENT_CALL_ENDED,
				data,
			)

			BroadcastToRoom(
				data["fromRoomId"].(string),
				EVENT_CALL_ENDED,
				data,
			)
		}
	})
}
