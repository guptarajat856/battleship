package model

import (
	"fmt"
)

type Player struct {
	Board Board
	Id    int
}

func (p *Player) AreAllItemsKilled() bool {
	return p.Board.AreAllItemsKilled()
}

func (p *Player) TakeHit(coordinate Coordinate) {
	err := p.Board.TakeHit(coordinate)
	if err != nil {
		fmt.Printf("%v", err)
	}
}
