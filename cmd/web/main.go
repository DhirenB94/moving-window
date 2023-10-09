package main

import (
	"log"
	"net/http"

	movingwindow "dhiren.brahmbhatt/moving-window"
)

func main() {
	store := movingwindow.NewInMemDB()
	server := movingwindow.NewRequestServer(store)
	err := http.ListenAndServe(":5001", server)
	if err != nil {
		log.Fatal(err)
	}
}
