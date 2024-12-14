package main

import (
	"chatapp/internal/chatapp/api"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", api.IndexHandler)
	http.HandleFunc("/ws", api.WsHandler)

	fmt.Println("Server started on port 6969")
	log.Fatal(http.ListenAndServe(":6969", nil))
}
