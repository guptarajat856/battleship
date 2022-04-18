package main

import (
	output "battleship/io/output"
	model "battleship/model"
	strategy "battleship/strategy"
	"strconv"
)

type GameLoop struct {
	Players                  [](*model.Player)
	WinnerStrategy           strategy.IWinnerStrategy
	Printer                  output.IOutputPrinter
	NextPlayerStrategy       strategy.IPlayerPickingStrategy
	ChanceGenerationStrategy strategy.IChanceGenerationStrategy
}

func NewGameLoop(players [](*model.Player), winnerStrategy strategy.IWinnerStrategy, printer output.IOutputPrinter,
	nextPlayerStrategy strategy.IPlayerPickingStrategy, chanceGenerationStrategy strategy.IChanceGenerationStrategy) *GameLoop {
	return &GameLoop{players, winnerStrategy, printer, nextPlayerStrategy, chanceGenerationStrategy}
}

func (g *GameLoop) TakeChance(playerObj *model.Player) model.PlayerChanceTarget {
	opponents := [](*model.Player){}
	for _, opponent := range g.Players {
		if opponent.Id != playerObj.Id {
			opponents = append(opponents, opponent)
		}
	}

	return g.ChanceGenerationStrategy.GetPlayerChanceTarget(opponents)
}

func (g *GameLoop) Start() {
	currentPlayerIndex := g.NextPlayerStrategy.FirstPlayer(g.Players)
	g.Printer.PrintMsg("Starting game!")
	for {
		currentPlayer := g.Players[currentPlayerIndex]

		g.Printer.PrintMsg("\n\nPlayer: " + strconv.Itoa(currentPlayer.Id) + " chance:")
		playerChanceTarget := g.TakeChance(currentPlayer)

		playerChanceTarget.PlayerO.TakeHit(playerChanceTarget.Target)

		g.Printer.PrintSelfBoard(*currentPlayer)
		g.Printer.PrintOpponentBoard(g.Players, *currentPlayer)

		winner := g.WinnerStrategy.GetWinner(g.Players)
		if winner != nil {
			g.Printer.PrintWinner(*winner)
			break
		}

		currentPlayerIndex = g.NextPlayerStrategy.PickNextPlayer(currentPlayerIndex, g.Players)
	}
}
