package main

type Player struct {
	x, y int
}

func (*Player) getPlaceholder() string {
	return "P"
}

func (p *Player) setCoords(newX, newY int) {
	p.y = newY
	p.x = newX
}

// Returns y and x
func (p *Player) getCoords() (int, int) {
	return p.y, p.x
}
