package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lanran2001/lanrand/handler/ws"
)

func main() {
	r := gin.Default()
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	r.GET("/", ws.EchoMessage)
	r.GET("/ws/echo_display", ws.DisplayEcho)
	err := r.Run(":8080")
	if err != nil {
		log.Println(err)
		return
	}
}
