package socket

import (
	"log"
	"video-conf/database/repositories"

	socketio "github.com/googollee/go-socket.io"
)

func InitBaseHandlers() {
	socket.OnConnect("/", func(s socketio.Conn) error {
		log.Println("client connected by ID:", s.ID())

		// join client to general room
		s.Join(SOCKET_GENERAL_ROOM)

		accessToken := getUserTkn(s.URL().RawQuery)

		if accessToken != "" {
			tokenRepo := repositories.NewTokenRepository()
			token := tokenRepo.FindByAccess(accessToken)

			s.SetContext(SocketContext{
				User: SocketUser{
					ID:    token.User.ID,
					Name:  token.User.Name,
					Email: token.User.Email,
				},
			})
		}

		return nil
	})

	socket.OnError("/", func(s socketio.Conn, e error) {
		log.Println("Socket OnError:", e.Error())
	})

	socket.OnDisconnect("/", func(s socketio.Conn, msg string) {
		log.Println("client disconnected by ID:", msg)
	})

	socket.OnEvent("/", EVENT_USER_GET, func(s socketio.Conn) {
		s.Emit(EVENT_USER_ME, SuccessResponse(s.Context()))
	})
}
