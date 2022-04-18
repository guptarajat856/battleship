package model

type IBoundary interface {
	Contains(Coordinate) bool
	AllCoordinates() []Coordinate
}
