package output

import (
	model "battleship/model"
	"fmt"
	"strconv"
)

type SysOutputPrinter struct {
}

func (s *SysOutputPrinter) PrintMsg(msg string) {
	fmt.Println(msg)
}

func (s *SysOutputPrinter) PrintWinner(playerObj model.Player) {
	fmt.Println("Game Finished!")
	fmt.Println("Player: " + strconv.Itoa(playerObj.Id) + " won")
}

func (s *SysOutputPrinter) playerInfo(playerObj model.Player) {
	s.PrintMsg("Player: " + strconv.Itoa(playerObj.Id))
}

func (s *SysOutputPrinter) PrintSelfBoard(playerObj model.Player) {
	s.PrintMsg("Your board status: ")
	s.playerInfo(playerObj)
	fmt.Printf("Board boundary: %+v\n", playerObj.Board.Boundary)
	fmt.Printf("Ships: %+v\n", playerObj.Board.Items)
	fmt.Printf("Hit locations: %+v\n", playerObj.Board.HitLocations())
	fmt.Printf("Missed locations: %+v\n", playerObj.Board.MissLocations())
}

func (s *SysOutputPrinter) printOpponentBoard(playerObj model.Player) {
	s.PrintMsg("\nOpponent board status: ")
	s.playerInfo(playerObj)
	fmt.Printf("Board boundary: %+v\n", playerObj.Board.Boundary)
	fmt.Printf("Hit locations: %+v\n", playerObj.Board.HitLocations())
	fmt.Printf("Missed locations: %+v\n", playerObj.Board.MissLocations())
}

func (s *SysOutputPrinter) PrintOpponentBoard(allPlayers [](*model.Player), playerObj model.Player) {
	for _, opponentPlayerObj := range allPlayers {
		if opponentPlayerObj.Id != playerObj.Id {
			s.printOpponentBoard(*opponentPlayerObj)
		}
	}
}
