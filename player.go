package game

import "fmt"

type Player interface {
	Play() Spotter
}

type Person struct {
	Name  string
	Shape string
}

func NewPlayer(name, shape string) Player {
	return &Person{Name: name, Shape: shape}
}

func (p *Person) Play() Spotter {
	var x, y int
	fmt.Println("pleae enter the coordinates x,y")
	fmt.Scanln(&x)
	fmt.Scanln(&y)
	spot := NewSpot(p.shape, p.name, x, y)
	return spot
}
