package game

type turn struct {
	listOfPlayer []Player
}

type GetPlayer func() Player

type Turner interface {
	Iterator() GetPlayer
	AddPlayer(p Player)
}

func NewTurner() Turner {
	return turn{}
}

func (t turn) AddPlayer(p player) {
	t.listOfPlayer = append(listOfPlayer, p)
}

func NewTwoPlayerTurner(p1, p2 Player) Turner {
	t := turn{}
	t.AddPlayer(p1)
	t.AddPlayer(p2)
	return t
}

func (t turn) Iterator() GetPlayer {
	index := 0
	mod := len(t.listOfPlayer)
	return func() Player {
		player := t.listOfPlayer[index]
		index = (index + 1) % mod
		return player
	}
}
