package model

type RectangularBoundary struct {
	TopLeft     Coordinate
	BottomRight Coordinate
}

func (rec *RectangularBoundary) Contains(coordinate Coordinate) bool {
	return (coordinate.X >= rec.TopLeft.X && coordinate.X <= rec.BottomRight.X && coordinate.Y <= rec.TopLeft.Y && coordinate.Y >= rec.BottomRight.Y)
}

func (rec *RectangularBoundary) AllCoordinates() []Coordinate {
	coordinates := []Coordinate{}
	for i := rec.TopLeft.X; i <= rec.BottomRight.X; i++ {
		for j := rec.BottomRight.Y; j <= rec.TopLeft.Y; j++ {
			coordinates = append(coordinates, Coordinate{i, j})
		}
	}

	return coordinates
}
