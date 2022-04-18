package main

import (
	input "battleship/io/input"
	output "battleship/io/output"
	model "battleship/model"
	strategy "battleship/strategy"
)

func main() {
	inputProvider := input.SysInputProvider{}
	players := [](*model.Player){}

	player1 := getPlayer1()
	players = append(players, &player1)
	player2 := getPlayer2()
	players = append(players, &player2)

	gameLoop := NewGameLoop(players, &(strategy.DefaultWinnerStrategy{}), &(output.SysOutputPrinter{}),
		&(strategy.RoundRobinPlayerPickingStrategy{}), &(strategy.LocalChanceGeneration{&inputProvider}))
	gameLoop.Start()
}

func getPlayer1() model.Player {
	rectangleBoundary := model.RectangularBoundary{model.Coordinate{0, 10}, model.Coordinate{10, 0}}

	ship1 := model.BoardItem{"Carrier", &(model.RectangularBoundary{model.Coordinate{0, 7}, model.Coordinate{4, 7}})}
	ship2 := model.BoardItem{"Battleship", &(model.RectangularBoundary{model.Coordinate{0, 7}, model.Coordinate{4, 7}})}
	ship3 := model.BoardItem{"Cruiser", &(model.RectangularBoundary{model.Coordinate{7, 3}, model.Coordinate{7, 5}})}
	ship4 := model.BoardItem{"Destroyer", &(model.RectangularBoundary{model.Coordinate{4, 9}, model.Coordinate{6, 9}})}
	ship5 := model.BoardItem{"Submarine", &(model.RectangularBoundary{model.Coordinate{3, 6}, model.Coordinate{4, 3}})}

	ships := []model.BoardItem{}
	ships = append(ships, ship1)
	ships = append(ships, ship2)
	ships = append(ships, ship3)
	ships = append(ships, ship4)
	ships = append(ships, ship5)

	board := model.NewBoard(ships, &rectangleBoundary)
	playerObj := model.Player{*board, 1}
	return playerObj
}

func getPlayer2() model.Player {
	rectangleBoundary := model.RectangularBoundary{model.Coordinate{0, 10}, model.Coordinate{10, 0}}

	ship1 := model.BoardItem{"Carrier", &(model.RectangularBoundary{model.Coordinate{0, 7}, model.Coordinate{4, 7}})}
	ship2 := model.BoardItem{"Battleship", &(model.RectangularBoundary{model.Coordinate{0, 7}, model.Coordinate{4, 7}})}
	ship3 := model.BoardItem{"Cruiser", &(model.RectangularBoundary{model.Coordinate{7, 3}, model.Coordinate{7, 5}})}
	ship4 := model.BoardItem{"Destroyer", &(model.RectangularBoundary{model.Coordinate{4, 9}, model.Coordinate{6, 9}})}
	ship5 := model.BoardItem{"Submarine", &(model.RectangularBoundary{model.Coordinate{3, 6}, model.Coordinate{4, 3}})}

	ships := []model.BoardItem{}
	ships = append(ships, ship1)
	ships = append(ships, ship2)
	ships = append(ships, ship3)
	ships = append(ships, ship4)
	ships = append(ships, ship5)

	board := model.NewBoard(ships, &rectangleBoundary)
	playerObj := model.Player{*board, 2}
	return playerObj
}
