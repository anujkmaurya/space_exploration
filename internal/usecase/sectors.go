package usecase

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/personal-work/space_exploration/internal/models"
)

//CreateSector : creates sector
func CreateSector(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	sector := &models.Sector{}
	json.NewDecoder(r.Body).Decode(sector)

	log.Println("sector ", sector)

	//check if sector is already created
	if _, ok := models.SectorsMap[sector.ID]; ok {
		return nil, createAppError("sector already exist", http.StatusConflict)
	}

	models.SectorIDCounter++

	//set sector ID
	sector.ID = models.SectorIDCounter

	//insert in global map
	models.SectorsMap[sector.ID] = sector

	return sector, nil
}

//GetAllSectors : get all sector info
func GetAllSectors(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	sectorList := []*models.Sector{}

	for _, sector := range models.SectorsMap {
		sectorList = append(sectorList, sector)
	}

	return sectorList, nil
}
