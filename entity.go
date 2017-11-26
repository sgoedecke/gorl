package main

import (
	"github.com/nsf/termbox-go"
)

// DynamicEntity type: Screens and worlds have these. They must be able to handle collision,
// check for collision, and know how to draw themselves

type DynamicEntity interface {
	HandleCollision(e *Entity, l *Log)
	Draw(x int, y int)
	CheckCollision(e *Entity, x int, y int) bool
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

func (target Entity) HandleCollision(e *Entity, l *Log) {
	l.AddMessage("You bumped into somebody", e.Color)
	l.AddMessage("Hey, don't bump into me!", target.Color)
}

func (entity Entity) CheckCollision(target *Entity, x int, y int) bool {
	if entity.X == x && entity.Y == y {
		entity.HandleCollision(target, entity.screen.World.Log)
		return true
	}
	return false
}

func (entity Entity) Draw(x int, y int) {
	termbox.SetCell(entity.X+x, entity.Y+y, entity.img, entity.Color, termbox.ColorBlack)
}
