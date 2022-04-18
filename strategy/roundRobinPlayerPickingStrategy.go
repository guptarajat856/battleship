package strategy

import model "battleship/model"

type RoundRobinPlayerPickingStrategy struct {
}

func (r *RoundRobinPlayerPickingStrategy) FirstPlayer(allPlayers [](*model.Player)) int {
	return 0
}

func (r *RoundRobinPlayerPickingStrategy) PickNextPlayer(curentPlayerIndex int, allPlayers [](*model.Player)) int {
	return (curentPlayerIndex + 1) % len(allPlayers)
}
