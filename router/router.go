package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lanran2001/lanrand/handler/ws"
)

func RegisterRoutes(r *mux.Router) {
	// serve static file request
	fs := http.FileServer(http.Dir("assets/"))
	serveFileHandler := http.StripPrefix("/static/", fs)
	r.PathPrefix("/static/").Handler(serveFileHandler)

	wsRouter := r.PathPrefix("/ws").Subrouter()
	wsRouter.HandleFunc("/echo", ws.EchoMessage)
	wsRouter.HandleFunc("/echo_display", ws.DisplayEcho)
}
