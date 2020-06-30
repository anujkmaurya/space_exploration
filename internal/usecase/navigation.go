package usecase

import (
	"net/http"

	"github.com/personal-work/space_exploration/internal/models"
)

//HealthCheck to check health of service
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("tested OK\n"))
}

//findLocationCoreLogic to find the location of drone
func findLocationCoreLogic(drone *models.Drone) float64 {
	location := float64(drone.SectorID)*drone.Loc.SumOfCoordinates() + drone.Vel
	return location
}

//findLocationAdvancedCoreLogic to find the location of drone , more advanced logic than FindLocationCoreLogic
func findLocationAdvancedCoreLogic(drone *models.Drone) float64 {
	//add more advanced core logic here
	location := float64(drone.SectorID)*drone.Loc.SumOfCoordinates() + drone.Vel + 100
	return location
}
