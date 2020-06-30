package models

//Drone struct declaration
type Drone struct {
	ID       uint64      `json:"id"`
	Name     string      `json:"name"`
	Type     string      `json:"type"`
	SectorID uint64      `json:"sector_id"`
	Loc      Coordinates `json:"coordinates"`
	Vel      float64     `json:"velocity"`
}

//Coordinates coordinates in x, y,z direction
type Coordinates struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

var DronesMap map[uint64]*Drone

var DroneIDCounter uint64

//SumOfCoordinates : to get sum of coordination
func (c Coordinates) SumOfCoordinates() float64 {
	return c.X + c.Y + c.Z
}

//SetCoordinates : sets present coordinats of drone
func (d *Drone) SetCoordinates(x, y, z float64) {
	d.Loc.X = x
	d.Loc.Y = y
	d.Loc.Z = z
}

//SetVelocity : sets present velocity of drone
func (d *Drone) SetVelocity(v float64) {
	d.Vel = v
}
