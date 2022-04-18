package model

type Coordinate struct {
	X int
	Y int
}

func (c *Coordinate) PresentInSlice(s []Coordinate) bool {
	for _, a := range s {
		if a.X == c.X && a.Y == c.Y {
			return true
		}
	}
	return false
}
