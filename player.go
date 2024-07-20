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
	fmt.Printf("%s pleae enter the coordinates x,y\n", p.Name)
	fmt.Scanln(&x)
	fmt.Scanln(&y)
	spot := NewSpot(p.Shape, p.Name, x, y)
	return spot
}
