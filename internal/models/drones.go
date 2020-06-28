package models

//Drone struct declaration
type Drone struct {
	ID   int64       `json:"id"`
	Name string      `json:"name"`
	Loc  Coordinates `json:"coordinates"`
	Vel  float64     `json:"velocity"`
}

//Location
type Coordinates struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type Location struct {
	Loc float64 `json:"loc"`
}

type Location2 struct {
	Loc float64 `json:"location"`
}
