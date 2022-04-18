package model

type BoardItem struct {
	Name     string
	Boundary IBoundary
}

func (b *BoardItem) IsKilled(hitLocations []Coordinate) bool {
	itemCoordinates := b.Boundary.AllCoordinates()
	for _, itemCoordinate := range itemCoordinates {
		if !itemCoordinate.PresentInSlice(hitLocations) {
			return false
		}
	}

	return true
}

func (b *BoardItem) ContainsCoordinate(coordinate Coordinate) bool {
	return b.Boundary.Contains(coordinate)
}
