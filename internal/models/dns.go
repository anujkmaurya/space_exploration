package models

//DNS struct declaration
type DNS struct {
	ID         uint64   `json:"id"`
	Name       string   `json:"name"`
	SectorList []uint64 `json:"sectors"`
}

var DNSMap map[uint64]*DNS

var DNSIDCounter uint64
