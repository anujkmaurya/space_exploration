package delivery

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/personal-work/space_exploration/internal/models"
)

func (d *Delivery) parseGetLocationRequest(w http.ResponseWriter, r *http.Request) (*models.LocationReq, error) {

	locReq := models.LocationReq{}
	json.NewDecoder(r.Body).Decode(&locReq)

	//get drone id
	params := mux.Vars(r)
	var id = params["id"]

	droneID, err := strconv.ParseUint(id, 0, 64)
	if err != nil {
		return nil, models.CreateAppError("droneID is not an integer", http.StatusBadRequest)
	}
	locReq.DroneID = droneID

	//get dns id
	var dnsIDStr = params["dnsID"]

	dnsID, err := strconv.ParseUint(dnsIDStr, 0, 64)
	if err != nil {
		return nil, models.CreateAppError("dnsID is not an integer", http.StatusBadRequest)
	}
	locReq.DnsID = dnsID

	//check typeof resp needed
	r.ParseForm()
	isCustomStr := strings.ToLower(strings.TrimSpace(r.FormValue("is_custom")))
	if isCustomStr == "true" {
		locReq.IsCustom = true
	}

	return &locReq, nil
}

//GetDroneLocation : get drone location
func (d *Delivery) GetDroneLocation(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	//parse function in delivery layer
	locReq, err := d.parseGetLocationRequest(w, r)
	if err != nil {
		log.Println("error occured in parsing req:", err.Error())
		return nil, err
	}

	resp, err := d.UsecaseLayer.GetDroneLocation(locReq)
	if err != nil {
		log.Println("error occured in getting drone location:", err.Error())
		return nil, err
	}

	return resp, nil
}
