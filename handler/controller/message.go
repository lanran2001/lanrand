package controller

import (
	"encoding/json"
	"log"
)

type serverRequest struct {
	Method string      `json:"method"`
	Data   interface{} `json:"data"`
}

func sendMessageToAllClients(msg []byte) {
	clients := GetClients()
	for _, client := range clients {
		if err := client.WriteMessage(1, msg); err != nil {
			log.Println(err)
		}
	}
}

func sendRequestToAllClients(req serverRequest) {
	msg, _ := json.Marshal(req)
	log.Println(req)
	sendMessageToAllClients(msg)
}
