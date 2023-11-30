package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"heji-server/internal/config"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
	return true
}} // use default options
var wsConn *websocket.Conn

var wsTimeout = 90 * time.Second

func WebSocket(router *gin.RouterGroup, config *config.Config) {
	router.GET("/ws", func(c *gin.Context) {
		w := c.Writer
		r := c.Request
		wsConn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		//var writeMutex sync.Mutex

		defer wsConn.Close()
		//connId := pkg.UUID()
		// 监听关闭
		wsConn.SetCloseHandler(func(code int, text string) error {
			println("conn closed")
			wsConn.CloseHandler()
			wsConn.Close()
			err = errors.New(text)
			return err
		})
		//go wsWriter(ws, &writeMutex, connId)
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
func SendWsMsg() {

}
