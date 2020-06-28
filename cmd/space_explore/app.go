package main

import (
	"log"
	"net/http"

	"github.com/personal-work/space_exploration/internal/delivery"
)

func main() {

	delivery.Init()

	// serve
	log.Printf("Server up on port 9001")
	log.Fatal(http.ListenAndServe(":9001", nil))
}
