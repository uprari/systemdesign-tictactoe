package game

type Spot struct {
	Shape      string
	PlayerName string
	X          int
	Y          int
}

func NewSpot(shape, name string, x, y int) Spotter {
	return &Spot{Shape: shape, PlayerName: name, X: x, Y: y}
}

func (s *Spot) GetShape() string {
	return s.Shape
}

func (s *Spot) GetPlayerName() string {
	return s.PlayerName
}

func (s *Spot) GetPosition() (int, int) {
	return s.X, s.Y
}
