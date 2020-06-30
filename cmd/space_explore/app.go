package main

import (
	"log"
	"net/http"

	"github.com/personal-work/space_exploration/internal/delivery"
	"github.com/personal-work/space_exploration/internal/usecase"
)

func main() {

	iUsecase := usecase.Init()
	delivery.Init(iUsecase)

	// serve
	log.Printf("Server up on port 9001")
	log.Fatal(http.ListenAndServe(":9001", nil))
}
