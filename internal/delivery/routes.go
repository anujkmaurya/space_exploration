package delivery

import (
	"github.com/personal-work/space_exploration/internal/usecase"

	"github.com/gorilla/mux"
)

//InitHandlers inits all the routes
func InitHandlers() *mux.Router {

	r := mux.NewRouter()

	//healthcheck api
	r.HandleFunc("/healthcheck", usecase.HealthCheck).Methods("GET")

	//Create sector
	r.Handle("/sectors", HandlerFunc(usecase.CreateSector)).Methods("POST")
	//Create sector
	r.Handle("/drones", HandlerFunc(usecase.CreateDrone)).Methods("POST")
	//Create sector
	r.Handle("/dns", HandlerFunc(usecase.CreateDNS)).Methods("POST")

	return r
}
