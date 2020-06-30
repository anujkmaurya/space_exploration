package usecase

import (
	"github.com/personal-work/space_exploration/internal/models"
)

func Init() UsecaseInterface {

	//init resources
	initResources()

	usecase := &Usecase{}
	return usecase
}

func initResources() {

	//init drone management map
	models.DronesMap = make(map[uint64]*models.Drone)

	//init Sector map
	models.SectorsMap = make(map[uint64]*models.Sector)

	//init DNS map
	models.DNSMap = make(map[uint64]*models.DNS)
}
