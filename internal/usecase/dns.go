package usecase

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/personal-work/space_exploration/internal/models"
)

//CreateDNS : creates new DNS
func (u *Usecase) CreateDNS(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	dns := &models.DNS{}
	json.NewDecoder(r.Body).Decode(dns)

	// log.Println("dns ", dns)

	//check if drone is already created
	if _, ok := models.DNSMap[dns.ID]; ok {
		return nil, models.CreateAppError("dns system already exist", http.StatusConflict)
	}

	for _, sectorID := range dns.SectorList {
		//check if sector exists
		if _, ok := models.SectorsMap[sectorID]; !ok {
			return nil, models.CreateAppError(fmt.Sprintf("sectorID:%d doesn't exist", sectorID), http.StatusBadRequest)
		}
	}

	//set Dns ID
	models.DNSIDCounter++
	dns.ID = models.DNSIDCounter

	//insert in global map
	models.DNSMap[dns.ID] = dns

	return dns, nil
}

//GetAllDNS : get all DNS info
func (u *Usecase) GetAllDNS(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	dnsList := []*models.DNS{}

	for _, dns := range models.DNSMap {
		dnsList = append(dnsList, dns)
	}

	return dnsList, nil
}

//GetDroneLocation : get drone location
func (u *Usecase) GetDroneLocation(droneLocReq *models.LocationReq) (interface{}, error) {

	droneID := droneLocReq.DroneID
	if _, ok := models.DronesMap[droneID]; !ok {
		return nil, models.CreateAppError("drone with this droneID doesn't exist", http.StatusBadRequest)
	}

	drone := models.DronesMap[droneID]

	//get dns info
	if _, ok := models.DNSMap[droneLocReq.DnsID]; !ok {
		return nil, models.CreateAppError("dns with this dnsID doesn't exist", http.StatusBadRequest)
	}

	dns := models.DNSMap[droneLocReq.DnsID]
	found := false
	for _, secID := range dns.SectorList {

		//check if sector of drone is monitored by given DNS
		if drone.SectorID == secID {
			found = true
			break
		}
	}
	if !found {
		return nil, models.CreateAppError("the DNS is not serving the sector in which Drone is located", http.StatusBadRequest)
	}

	//set present coordinates and velocity of drone
	drone.SetCoordinates(droneLocReq.X, droneLocReq.Y, droneLocReq.Z)
	drone.SetVelocity(droneLocReq.Vel)

	return u.findLocationReturnType(drone, droneLocReq.IsCustom), nil
}
