package models

//Sector struct declaration
type Sector struct {
	ID   uint64 `json:"sector_id"`
	Name string `jso:"sector_name"`
}

var SectorsMap map[uint64]*Sector

var SectorIDCounter uint64
