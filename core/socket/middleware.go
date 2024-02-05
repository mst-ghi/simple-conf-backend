package socket

import (
	"reflect"
	"strings"

	socketio "github.com/googollee/go-socket.io"
)

type HandlerT interface{}

func create(middleware interface{}) func(handler interface{}) HandlerT {
	mw := reflect.ValueOf(middleware)
	mwT := mw.Type()
	mwNumIn := mwT.NumIn()
	nextT := mwT.In(mwNumIn - 1)

	return func(handler interface{}) HandlerT {
		h := reflect.ValueOf(handler)
		hT := h.Type()

		wrapper := reflect.MakeFunc(hT, func(in []reflect.Value) []reflect.Value {
			next := reflect.MakeFunc(nextT, func([]reflect.Value) []reflect.Value {
				res := reflect.ValueOf(handler).Call(in)
				return res[:nextT.NumOut()]
			})

			mwParams := make([]reflect.Value, mwNumIn-1)
			for i := range mwParams {
				if i < len(in) {
					mwParams[i] = in[i]
				}
			}
			mwParams = append(mwParams, next)

			return mw.Call(mwParams)
		})

		return wrapper.Interface().(HandlerT)
	}
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

var SocketAuth = create(func(s socketio.Conn, next func()) {
	token := getUserTkn(s.URL().RawQuery)

	if token == "" {
		s.Emit("unauthorized")
	}

	s.SetContext(map[string]string{
		"token": token,
	})

	next()
})
