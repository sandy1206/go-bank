package main

import (
	"fmt"
	"log"
)

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		// Handle the error here
		// For example, you can log the error or return an error response
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", store)
	server := NewAPIServer(":3000", store)
	server.Run()
}
