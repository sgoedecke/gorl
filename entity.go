package main

import (
	"github.com/nsf/termbox-go"
)

// Entity type: has coords, a rune, and knows what world it's in. Can move around.
// Drawing is handled by the World.

type Entity struct {
	X      int
	Y      int
	img    rune
	Health int
	Color  termbox.Attribute
	screen  *Screen
}

type DynamicEntity interface {
	HandleCollision(e *Entity, l *Log)
	Draw(x int, y int)
	CheckCollision(e *Entity, x int, y int) bool
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

// HealthBar type: has a Entity, knows how to draw itself

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

// Portal type: colliding with it teleports the player to a new screen

type Portal struct {
	Entity
	Destination *Screen
	destX int
	destY int
}

func (p Portal) HandleCollision(e *Entity, l *Log) {
	l.AddMessage("You passed into a new area", termbox.ColorWhite)
	player := e
	w := player.screen.World
	w.ActiveScreen = p.Destination
	player.X = p.destX
	player.Y = p.destY
}
