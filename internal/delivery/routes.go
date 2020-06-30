package delivery

import (
	"github.com/personal-work/space_exploration/internal/usecase"

	"github.com/gorilla/mux"
)

//InitHandlers inits all the routes
func (d *Delivery) InitHandlers() *mux.Router {

	r := mux.NewRouter()

	//healthcheck api
	r.HandleFunc("/healthcheck", usecase.HealthCheck).Methods("GET")

	//sector api's
	r.Handle("/sectors", HandlerFunc(d.UsecaseLayer.CreateSector)).Methods("POST")
	r.Handle("/sectors", HandlerFunc(d.UsecaseLayer.GetAllSectors)).Methods("GET")

	//drone api's
	r.Handle("/drones", HandlerFunc(d.UsecaseLayer.CreateDrone)).Methods("POST")
	r.Handle("/drones", HandlerFunc(d.UsecaseLayer.GetAllDrones)).Methods("GET")

	//Dns api's
	r.Handle("/dns", HandlerFunc(d.UsecaseLayer.CreateDNS)).Methods("POST")
	r.Handle("/dns", HandlerFunc(d.UsecaseLayer.GetAllDNS)).Methods("GET")
	//api to find location of given drone w.r.t given dns
	r.Handle("/dns/{dnsID}/drones/{id}/location", HandlerFunc(d.GetDroneLocation)).Methods("POST")

	return r
}
