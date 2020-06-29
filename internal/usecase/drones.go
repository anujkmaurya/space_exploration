package usecase

import (
	"encoding/json"
	"net/http"

	"github.com/personal-work/space_exploration/internal/models"
)

// //Drone struct declaration
// type Drone struct {
// 	ID       uint64      `json:"id"`
// 	Name     string      `json:"name"`
// 	SectorID uint64      `json:"sector_id"`
// 	Loc      Coordinates `json:"coordinates"`
// 	Vel      float64     `json:"velocity"`
// }

// //Coordinates coordinates in x, y,z direction
// type Coordinates struct {
// 	X float64 `json:"x"`
// 	Y float64 `json:"y"`
// 	Z float64 `json:"z"`
// }

// //Location location of drone
// type Location struct {
// 	Loc float64 `json:"loc"`
// }

// //Location2 custom response for location of drone
// type Location2 struct {
// 	Loc float64 `json:"location"`
// }

// var DronesMap map[uint64]*Drone

// var DroneIDCounter uint64

//CreateDrone : creates new drone
func CreateDrone(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	drone := &models.Drone{}
	json.NewDecoder(r.Body).Decode(drone)

	// log.Println("drone ", drone)

	//check if drone is already created
	if _, ok := models.DronesMap[drone.ID]; ok {
		return nil, createAppError("drone already exist", http.StatusConflict)
	}

	models.DroneIDCounter++
	//set drone ID
	drone.ID = models.SectorIDCounter

	//insert in global map
	models.DronesMap[drone.ID] = drone

	return drone, nil
}
