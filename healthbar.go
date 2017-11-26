package main

import (
	"github.com/nsf/termbox-go"
)

// HealthBar type: has a Player, knows how to draw itself

type HealthBar struct {
	Player *Player
}

func (h *HealthBar) Draw(x int, y int) {
	width := 80 // world width
	w := float32(h.Player.Health) * (float32(width) / 100.0)
	for xIndex := x; xIndex < int(w)+x; xIndex++ {
		termbox.SetCell(xIndex, y, 35, termbox.ColorRed, termbox.ColorBlack)
	}
}
