package models

//Location location of drone
type Location struct {
	Loc float64 `json:"loc"`
}

//LocationCustom custom response for location of drone for friendly company
type LocationCustom struct {
	Loc float64 `json:"location"`
}

type LocationReq struct {
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	Z        float64 `json:"z"`
	Vel      float64 `json:"vel"`
	IsCustom bool
	DroneID  uint64
}
