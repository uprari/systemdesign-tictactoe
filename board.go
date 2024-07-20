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
	check := b.layout[x][y]
	if check != nil {
		fmt.Println("Sorry %s this spot is already taken by %s", s.GetPlayerName(), check.GetPlayerName())
		return
	}
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
	for j := 0; j < b.y; j++ {
		fmt.Print("__ ")
	}
	fmt.Print("\n")
	for i := 0; i < b.x; i++ {

		for j := 0; j < b.y; j++ {
			valueAtSpotter := b.GetValue(i, j)
			fmt.Printf("|%s", valueAtSpotter)

		}
		fmt.Print(" |")
		fmt.Print("\n")
		for j := 0; j < b.y; j++ {
			fmt.Print("__ ")
		}
		fmt.Print("\n")
	}
}

func (b *Board) GetValue(i, j int) string {
	if b.layout[i][j] == nil {
		return " "
	}
	return b.layout[i][j].GetShape()
}

func NewBoard(x, y int) *Board {
	b := &Board{
		x: x,
		y: y,
	}
	b.layout = make([][]Spotter, x)
	for i := range b.layout {
		b.layout[i] = make([]Spotter, y)
	}
	return b
}
