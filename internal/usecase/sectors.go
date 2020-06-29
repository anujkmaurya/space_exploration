package usecase

import (
	"encoding/json"
	"net/http"

	"github.com/personal-work/space_exploration/internal/models"
)

//CreateSector : creates sector
func CreateSector(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	sector := &models.Sector{}
	json.NewDecoder(r.Body).Decode(sector)

	// log.Println("sector ", sector)

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
