package routes

import (
	"github.com/personal-work/space_exploration/internal/usecase"

	"github.com/gorilla/mux"
)

//InitHandlers inits all the routes
func InitHandlers() *mux.Router {

	r := mux.NewRouter()

	//healthcheck api
	r.HandleFunc("/healthcheck", usecase.HealthCheck).Methods("GET")
	return r
}
