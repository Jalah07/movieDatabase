package main

import (
	"log"
	"movieDatabase/handlers"
)

func main() {
	server := handlers.NewServer()

	log.Fatal(server.ListenAndServe())

}
