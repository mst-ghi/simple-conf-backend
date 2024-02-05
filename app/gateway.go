package app

import (
	"log"
	"video-conf/core/socket"

	socketio "github.com/googollee/go-socket.io"
)

func RegisterGateway() {
	io := socket.GetSocket()

	io.OnEvent("/", socket.EVENT_USER_GET, func(s socketio.Conn) {
		log.Println(socket.EVENT_USER_GET)

		s.Emit(socket.EVENT_USER_ME, s.Context())
	})
}
