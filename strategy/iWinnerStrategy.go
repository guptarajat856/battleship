package strategy

import model "battleship/model"

type IWinnerStrategy interface {
	GetWinner(playerList [](*model.Player)) *model.Player
}
