package controller

import "math/rand"

func draw() {
	req := serverRequest{
		Method: "draw",
		Data:   rand.Intn(100),
	}
	sendRequestToAllClients(req)
}
