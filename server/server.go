package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Start ...
func Start() {

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Login")
	})

	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Server run...")
	log.Fatal(s.ListenAndServe())

}
