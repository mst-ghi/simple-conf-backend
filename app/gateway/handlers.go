package gateway

import (
	"log"
	"strings"
	"video-conf/database/repositories"

	socketio "github.com/googollee/go-socket.io"
)

func ModuleHandlers() {
	socket.OnEvent("/", EVENT_USER_GET, func(s socketio.Conn) {
		s.Emit(EVENT_USER_ME, SuccessResponse(s.Context()))
	})
}

func BaseHandlers() {
	socket.OnConnect("/", func(s socketio.Conn) error {
		log.Println("Socket connected by ID:", s.ID())

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
		log.Println("Socket disconnected by ID:", msg)
	})
}

func getUserTkn(query string) string {
	tkn := strings.Split(query, "&")

	if tkn[0] != "" {
		for i := 0; i < len(tkn); i++ {
			if strings.Contains(tkn[i], "tkn=") {
				return strings.Split(tkn[i], "=")[1]
			}
		}
	}

	return ""
}
