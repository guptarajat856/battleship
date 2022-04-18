package output

import (
	model "battleship/model"
)

type IOutputPrinter interface {
	PrintMsg(msg string)
	PrintWinner(playerObj model.Player)
	PrintSelfBoard(playerObj model.Player)
	PrintOpponentBoard(allPlayers [](*model.Player), playerobj model.Player)
}
