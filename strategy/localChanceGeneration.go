package strategy

import (
	input "battleship/io/input"
	model "battleship/model"
)

type LocalChanceGeneration struct {
	InputProvider input.IInputProvider
}

func (l *LocalChanceGeneration) GetPlayerChanceTarget(opponents [](*model.Player)) model.PlayerChanceTarget {
	var targetPlayer *model.Player
	playerInput := l.InputProvider.TakeInput()
	for _, playerObj := range opponents {
		if playerObj.Id == playerInput.PlayerID {
			targetPlayer = playerObj
		}
	}

	attackCoordinate := model.Coordinate{playerInput.TargetX, playerInput.TargetY}
	return model.PlayerChanceTarget{attackCoordinate, targetPlayer}
}
