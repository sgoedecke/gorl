package main

import (
	"github.com/nsf/termbox-go"
)

// HealthBar type: has a Entity, knows how to draw itself
// This is the health bar for an entity (like the player)

type HealthBar struct {
	Entity *Entity
}

func (h *HealthBar) Draw(x int, y int) {
	width := 80 // world width
	w := float32(h.Entity.Health) * (float32(width) / 100.0)
	for xIndex := x; xIndex < int(w)+x; xIndex++ {
		termbox.SetCell(xIndex, y, 35, termbox.ColorRed, termbox.ColorBlack)
	}
}
