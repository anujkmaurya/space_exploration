package models

//DNS struct declaration
type DNS struct {
	ID       int64    `json:"id"`
	Name     string   `json:"name"`
	SectorID []Sector `json:"sectors"`
}

//SectorList
type SectorList struct {
	Sectors []Sector `json:"sectors"`
}
