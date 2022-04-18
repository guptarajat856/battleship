package strategy

import model "battleship/model"

type DefaultWinnerStrategy struct {
}

func (d *DefaultWinnerStrategy) GetWinner(playerList [](*model.Player)) *model.Player {
	alivePlayers := [](*model.Player){}
	for _, playerObj := range playerList {
		if !playerObj.AreAllItemsKilled() {
			alivePlayers = append(alivePlayers, playerObj)
		}
	}

	if len(alivePlayers) == 1 {
		return alivePlayers[0]
	}

	return nil
}
