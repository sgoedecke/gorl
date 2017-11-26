package main

import (
	"github.com/nsf/termbox-go"
)

// DynamicEntity type: Screens and worlds have these. They must be able to handle collision,
// check for collision, and know how to draw themselves

type DynamicEntity interface {
	CheckCollision(e DynamicEntity, x int, y int) bool // returns true if it's about to collide with e
	HandleCollision(e DynamicEntity)                   // takes an action on collision with e
	Draw(x int, y int)                           // draws self to the termbox buffer
	Act()                                        // does something per tick
}

// Entity type: has coords, a rune, and knows what world it's in. Can move around.
// Implements DynamicEntity. Other entity structs should mix in `Entity` and re-implement
// interface methods if they need to be overridden.

type Entity struct {
	X      int
	Y      int
	img    rune
	Health int
	Color  termbox.Attribute
	screen *Screen
}

func (e *Entity) MoveUp() {
	if e.screen.IsTileOccupied(e, e.X, e.Y-1) {
		return
	}
	e.Y--
}
func (e *Entity) MoveDown() {
	if e.screen.IsTileOccupied(e, e.X, e.Y+1) {
		return
	}
	e.Y++
}
func (e *Entity) MoveLeft() {
	if e.screen.IsTileOccupied(e, e.X-1, e.Y) {
		return
	}
	e.X--
}
func (e *Entity) MoveRight() {
	if e.screen.IsTileOccupied(e, e.X+1, e.Y) {
		return
	}
	e.X++
}

func (self Entity) Log() *Log {
	return self.screen.World.Log
}

// self has run into e
func (self *Entity) HandleCollision(e DynamicEntity) {
}

// target is trying to move into self
func (self *Entity) CheckCollision(target DynamicEntity, x int, y int) bool {
	if self.X == x && self.Y == y {
		target.HandleCollision(self)
		return true
	}
	return false
}

func (self *Entity) Draw(x int, y int) {
	termbox.SetCell(self.X+x, self.Y+y, self.img, self.Color, termbox.ColorBlack)
}

func (self *Entity) Act() {
	// do nothing
}
