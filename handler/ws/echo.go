package ws

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/lanran2001/lanrand/handler/controller"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func EchoMessage(c *gin.Context) {
	conn, _ := upgrader.Upgrade(c.Writer, c.Request, nil) // error ignored for sake of simplicity
	controller.AddClient(conn)
	for {
		// 读取客户端的消息
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		fmt.Printf("%s sent: %s, type: %d\n", conn.RemoteAddr(), string(msg), msgType)
		controller.Handle(msg)
		/*SendMessageToClient(conn.RemoteAddr().String(), string(msg), msgType)
		if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		}*/
	}
}

func SendMessageToClient(addr, msg string, msgType int) {
	clients := controller.GetClients()
	for _, client := range clients {
		if client.RemoteAddr().String() != addr {
			if err := client.WriteMessage(msgType, []byte(msg)); err != nil {
				log.Println(err)
			}
		}
	}
}
