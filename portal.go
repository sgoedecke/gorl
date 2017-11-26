package main

import (
	"github.com/nsf/termbox-go"
)

// Portal type: colliding with it teleports the player to a new screen
// Implements the DynamicEntity interface, uses the Entity mixin. Overrides
// collision-checking methods but borrows `Draw` from Entity

type Portal struct {
	Entity
	Destination *Screen
	destX       int
	destY       int
}

func (p Portal) HandleCollision(e *Entity, l *Log) {
	l.AddMessage("You passed into a new area", termbox.ColorWhite)
	player := e
	w := player.screen.World
	w.ActiveScreen = p.Destination
	player.screen = w.ActiveScreen
	player.X = p.destX
	player.Y = p.destY
}

func (portal Portal) CheckCollision(target *Entity, x int, y int) bool {
	if portal.X == x && portal.Y == y {
		portal.HandleCollision(target, portal.screen.World.Log)
		return true
	}
	return false
}
