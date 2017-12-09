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

func (self *Portal) HandleCollision(e DynamicEntity) {
	self.Log().AddMessage("You passed into a new area", termbox.ColorWhite)
	w := self.screen.World
	w.ActiveScreen = self.Destination
	w.Player.screen = w.ActiveScreen
	w.Player.X = self.destX
	w.Player.Y = self.destY
}

func (self *Portal) CollideWith(e DynamicEntity) {
}
