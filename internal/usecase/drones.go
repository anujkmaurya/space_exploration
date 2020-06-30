package usecase

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/personal-work/space_exploration/internal/models"
)

//CreateDrone : creates new drone
func (u *Usecase) CreateDrone(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	drone := &models.Drone{}
	json.NewDecoder(r.Body).Decode(drone)

	log.Println("drone ", drone)

	//check if drone is already created
	if _, ok := models.DronesMap[drone.ID]; ok {
		return nil, models.CreateAppError("drone already exist", http.StatusConflict)
	}

	//check if sector exists
	if _, ok := models.SectorsMap[drone.SectorID]; !ok {
		return nil, models.CreateAppError("sector doesn't exist", http.StatusBadRequest)
	}

	models.DroneIDCounter++
	//set drone ID
	drone.ID = models.DroneIDCounter

	if drone.Type == "" {
		drone.Type = "v1"
	}

	sector := models.SectorsMap[drone.SectorID]
	sector.DroneList = append(sector.DroneList, drone)

	//insert in global map
	models.DronesMap[drone.ID] = drone

	return drone, nil
}

//GetAllDrones : get all drones info
func (u *Usecase) GetAllDrones(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	droneList := []*models.Drone{}

	for _, drone := range models.DronesMap {
		droneList = append(droneList, drone)
	}

	return droneList, nil
}

//findLocation : wrapper to call desired corelogic as per drone type
func (u *Usecase) findLocation(drone *models.Drone) float64 {
	if drone.Type == "latest" {
		return u.findLocationAdvancedCoreLogic(drone)
	}
	return u.findLocationCoreLogic(drone)
}

//findLocationReturnType : wrapper to call desired corelogic as per drone type
func (u *Usecase) findLocationReturnType(drone *models.Drone, isCutom bool) interface{} {
	var normalResp models.Location
	var customResp models.LocationCustom

	if isCutom {
		customResp.Loc = u.findLocation(drone)
		return customResp
	}
	normalResp.Loc = u.findLocation(drone)
	return normalResp
}
