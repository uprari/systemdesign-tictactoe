package game

import "fmt"

type Spotter interface {
	GetPlayerName() string
	GetShape() string
	GetPosition() (int, int)
}

type Board struct {
	x          int
	y          int
	layout     [][]Spotter
	found      bool
	WinnerName string
}

func (b *Board) Mark(s Spotter) {
	x, y := s.GetPosition()
	b.layout[x][y] = s
}
func (b *Board) DeclareWinner(winnerName string) {
	b.found = true
	b.WinnerName = winnerName
}

func (b *Board) WinnerDecided() (bool, string) {
	return b.found, b.WinnerName
}

func (b *Board) Draw() {
	for i := 0; i < b.x; i++ {

		for j := 0; j < b.y; j++ {
			valueAtSpotter := b.GetValue(i, j)
			fmt.Println("%s", valueAtSpotter)
			fmt.Println(" |")
		}
		for j := 0; j < b.y; j++ {
			fmt.Println("_ ")
		}
	}
}

func (b *Board) GetValue(i, j int) string {
	if b.layout[i][j] == nil {
		""
	}
	return b.layout[i][j].GetShape()
}

func NewBoard(x, y int) *Board {
	b := &Board{
		x: x,
		y: y,
	}
	b.layout = make([][]Spotter, x)
	for _, val := range b.layout {
		val = make([]Spotter, y)
	}
	return b
}
