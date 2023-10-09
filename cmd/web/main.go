package main

import (
	"log"
	"net/http"

	movingwindow "dhiren.brahmbhatt/moving-window"
)

func main() {
	server := movingwindow.NewRequestServer(nil)
	err := http.ListenAndServe(":5001", server)
	if err != nil {
		log.Fatal(err)
	}
}
