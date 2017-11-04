package main

type Player struct {
	X     int
	Y     int
	img   rune
	world *World
}

func (p *Player) MoveUp() {
	if p.world.IsTileOccupied(p.X, p.Y-1) {
		return
	}
	p.Y--
}
func (p *Player) MoveDown() {
	if p.world.IsTileOccupied(p.X, p.Y+1) {
		return
	}
	p.Y++
}
func (p *Player) MoveLeft() {
	if p.world.IsTileOccupied(p.X-1, p.Y) {
		return
	}
	p.X--
}
func (p *Player) MoveRight() {
	if p.world.IsTileOccupied(p.X+1, p.Y) {
		return
	}
	p.X++
}
