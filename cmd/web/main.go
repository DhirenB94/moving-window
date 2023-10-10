package main

import (
	"log"
	"net/http"
	"os"

	movingwindow "dhiren.brahmbhatt/moving-window"
)

const fileName = "count.db.json"

func main() {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("unable to open file %s, %v", fileName, err)
	}

	store, err := movingwindow.NewFileSystem(file)
	if err != nil {
		log.Fatalf("unable to create file system %v", err)
	}
	server := movingwindow.NewRequestServer(store)

	err = http.ListenAndServe(":5001", server)
	if err != nil {
		log.Fatalf("could not listen on port 5001 %v", err)
	}
}
