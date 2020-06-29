package models

//Sector struct declaration
type Sector struct {
	ID        uint64   `json:"sector_id"`
	Name      string   `json:"sector_name"`
	DroneList []*Drone `json:"drone_list"`
}

var SectorsMap map[uint64]*Sector

var SectorIDCounter uint64
