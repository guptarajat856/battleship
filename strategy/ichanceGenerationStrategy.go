package strategy

import (
	model "battleship/model"
)

type IChanceGenerationStrategy interface {
	GetPlayerChanceTarget([](*model.Player)) model.PlayerChanceTarget
}
