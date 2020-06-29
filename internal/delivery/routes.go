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

	//sector api's
	r.Handle("/sectors", HandlerFunc(usecase.CreateSector)).Methods("POST")
	r.Handle("/sectors", HandlerFunc(usecase.GetAllSectors)).Methods("GET")

	//drone api's
	r.Handle("/drones", HandlerFunc(usecase.CreateDrone)).Methods("POST")
	r.Handle("/drones", HandlerFunc(usecase.GetAllDrones)).Methods("GET")

	//Dns api's
	r.Handle("/dns", HandlerFunc(usecase.CreateDNS)).Methods("POST")
	r.Handle("/dns", HandlerFunc(usecase.GetAllDNS)).Methods("GET")

	return r
}
