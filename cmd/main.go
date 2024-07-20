package main

import (
	"fmt"

	game "github.com/uprari/systemdesign-tictactoe"
)

func main() {
	board := game.NewBoard(3, 3)
	marker := game.BuildEvaluator(board, 3, board)
	playerone := game.NewPlayer("vaibahv", "@")
	playertwo := game.NewPlayer("rajesh", "*")
	winnerfound := false
	winnerName := ""
	turner := game.NewTwoPlayerTurner(playerone, playertwo)
	getPlayerWithTurn := turner.Iterator()
	for ; ; winnerfound, winnerName = board.WinnerDecided() {
		if winnerfound {
			break
		}
		p := getPlayerWithTurn()
		spot := p.Play()
		marker.Mark(spot)
		board.Draw()
	}
	fmt.Println("The winner is %v", winnerName)
}
