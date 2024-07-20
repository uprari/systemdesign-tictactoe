package main

import (
	"fmt"

	"github.com/uprari/systemdesign-tictactoe/game"
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
		p := getPlayerWithTurn()
		spot := p.play()
		marker(spot)
	}
	fmt.Println("The winner is %v", board)
}
