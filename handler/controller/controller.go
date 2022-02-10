package controller

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

var clients []*websocket.Conn

type clientRequest struct {
	Method string      `json:"method"`
	Data   interface{} `json:"data"`
}

func AddClient(client *websocket.Conn) {
	clients = append(clients, client)
}

func GetClients() []*websocket.Conn {
	return clients
}

func Handle(msg []byte) error {
	var req clientRequest
	err := json.Unmarshal(msg, &req)
	if err != nil {
		return err
	}
	switch req.Method {
	case "console":
		var cmd string
		cmd = req.Data.(string)
		draw()
		log.Println(cmd)
	}
	return nil
}
