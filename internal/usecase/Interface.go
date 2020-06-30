package usecase

import (
	"net/http"

	"github.com/personal-work/space_exploration/internal/models"
)

type UsecaseInterface interface {
	CreateDrone(w http.ResponseWriter, r *http.Request) (interface{}, error)
	GetAllDrones(w http.ResponseWriter, r *http.Request) (interface{}, error)

	CreateSector(w http.ResponseWriter, r *http.Request) (interface{}, error)
	GetAllSectors(w http.ResponseWriter, r *http.Request) (interface{}, error)

	//CreateDNS : creates new DNS
	CreateDNS(w http.ResponseWriter, r *http.Request) (interface{}, error)
	GetAllDNS(w http.ResponseWriter, r *http.Request) (interface{}, error)
	GetDroneLocation(droneLocReq *models.LocationReq) (interface{}, error)
}
