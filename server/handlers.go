package server

import (
	"net/http"

	"github.com/dm/chat-x-back/auth"
)

func initHandlers() {
	http.HandleFunc("/api/registr", auth.Registr)
}
