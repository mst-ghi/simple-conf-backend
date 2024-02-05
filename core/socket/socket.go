package socket

import (
	"log"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
)

var socket *socketio.Server

func GetSocket() *socketio.Server {
	return socket
}

func Initialize() {
	socket = socketio.NewServer(&engineio.Options{
		SessionIDGenerator: &CuidGenerator{},
	})

	InitBaseHandlers()
}

func Serve(engine *gin.Engine) {
	go func() {
		if err := socket.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()

	engine.GET("/socket.io/*any", gin.WrapH(socket))
	engine.POST("/socket.io/*any", gin.WrapH(socket))
}
