package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"heji-server/internal/config"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
	return true
}} // use default options
var wsc *websocket.Conn

func HandlerHoldWS(r *gin.RouterGroup, config *config.Config) {
	r.GET("/sync", func(c *gin.Context) {
		wsConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		wsc = wsConn
		defer wsc.Close()
		wsc.SetCloseHandler(func(code int, text string) error {
			println("conn closed")
			wsConn.CloseHandler()
			wsConn.Close()
			err = errors.New(text)
			return err
		})
		for {
			if err != nil {
				break
			}
			handlerMessage(wsConn)
		}
	})
}

func handlerMessage(wsConn *websocket.Conn) {
	mt, message, err := wsConn.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return
	}
	log.Printf("recv: %s", message)

	err = wsConn.WriteMessage(mt, message)
	if err != nil {
		log.Println("write:", err)
		return
	}
}
