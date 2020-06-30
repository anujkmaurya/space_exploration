package usecase

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/personal-work/space_exploration/internal/models"
)

// //DNS struct declaration
// type DNS struct {
// 	ID       uint64   `json:"id"`
// 	Name     string   `json:"name"`
// 	SectorID []Sector `json:"sectors"`
// }

// // //SectorList
// // type SectorList struct {
// // 	Sectors []Sector `json:"sectors"`
// // }

// var DNSMap map[uint64]*DNS

//CreateDNS : creates new DNS
func CreateDNS(w http.ResponseWriter, r *http.Request) (interface{}, error) {
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
func GetAllDNS(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	dnsList := []*models.DNS{}

	for _, dns := range models.DNSMap {
		dnsList = append(dnsList, dns)
	}

	return dnsList, nil
}
