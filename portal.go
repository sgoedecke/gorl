package main

import (
	"github.com/nsf/termbox-go"
)

// Portal type: colliding with it teleports the player to a new screen
// Implements the DynamicEntity interface, uses the Entity mixin. Overrides
// collision-checking methods but borrows `Draw` and `Act` from Entity

type Portal struct {
	Entity
	Destination *Screen
	destX       int
	destY       int
}

func (self *Portal) HandleCollision(e *Entity) {
	self.Log().AddMessage("You passed into a new area", termbox.ColorWhite)
	player := e
	w := player.screen.World
	w.ActiveScreen = self.Destination
	player.screen = w.ActiveScreen
	player.X = self.destX
	player.Y = self.destY
}

func (portal *Portal) CheckCollision(target *Entity, x int, y int) bool {
	if portal.X == x && portal.Y == y {
		portal.HandleCollision(target)
		return true
	}
	return false
}
