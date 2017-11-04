package main

import (
	"github.com/nsf/termbox-go"
)

// Player type: has coords, a rune, and knows what world it's in. Can move around.
// Drawing is handled by the World.

type Player struct {
	X      int
	Y      int
	img    rune
	Health int
	world  *World
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

// HealthBar type: has a Player, knows how to draw itself

type HealthBar struct {
	Player *Player
}

func (h *HealthBar) Draw(y int, width int) {
	w := float32(h.Player.Health) * (float32(width) / 100.0)
	for x := 0; x < int(w); x++ {
		termbox.SetCell(x, y, 35, termbox.ColorRed, termbox.ColorBlack)
	}
}
