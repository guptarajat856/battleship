package strategy

import model "battleship/model"

type IPlayerPickingStrategy interface {
	FirstPlayer(allPlayers [](*model.Player)) int
	PickNextPlayer(curentPlayerIndex int, allPlayers [](*model.Player)) int
}
