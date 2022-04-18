package model

import (
	"errors"
)

type Board struct {
	Items            []BoardItem
	Boundary         IBoundary
	AllBombLocations []Coordinate
}

func NewBoard(items []BoardItem, bound IBoundary) *Board {
	allBombLocations := []Coordinate{}
	board := Board{items, bound, allBombLocations}

	return &board
}

func (b *Board) TakeHit(coordinate Coordinate) error {
	if !b.Boundary.Contains(coordinate) {
		return errors.New("cooridnate out of Board's boundary")
	}

	b.AllBombLocations = append(b.AllBombLocations, coordinate)
	return nil
}

func (b *Board) AreAllItemsKilled() bool {
	for _, item := range b.Items {
		if !item.IsKilled(b.AllBombLocations) {
			return false
		}
	}

	return true
}

func (b *Board) HitLocations() []Coordinate {
	coordinates := []Coordinate{}
	for _, coordinate := range b.AllBombLocations {
		for _, item := range b.Items {
			if item.ContainsCoordinate(coordinate) {
				coordinates = append(coordinates, coordinate)
				break
			}
		}
	}

	return coordinates
}

func (b *Board) MissLocations() []Coordinate {
	coordinates := []Coordinate{}
	for _, coordinate := range b.AllBombLocations {
		doesBelongToShip := false
		for _, item := range b.Items {
			if item.ContainsCoordinate(coordinate) {
				doesBelongToShip = true
				break
			}
		}
		if !doesBelongToShip {
			coordinates = append(coordinates, coordinate)
		}
	}

	return coordinates
}
