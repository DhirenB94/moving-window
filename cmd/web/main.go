package main

import (
	"log"
	"net/http"

	movingwindow "dhiren.brahmbhatt/moving-window"
)

func main() {
	handler := http.HandlerFunc(movingwindow.RequestServer)
	err := http.ListenAndServe(":5001", handler)
	if err != nil {
		log.Fatal(err)
	}
}
