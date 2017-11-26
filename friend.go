package main

import (
	"math/rand"
	"time"
)

// Friend type: an entity that does not try to attack the player
// Implements the DynamicEntity interface, uses the Entity mixin. Overrides
// collision-checking methods but borrows `Draw` and `Act` from Entity

type Friend struct {
	Entity
}

func (self *Friend) HandleCollision(e *Entity) {
	self.Log().AddMessage("Hello, friend!", self.Color)
}

func (self *Friend) CheckCollision(target *Entity, x int, y int) bool {
	if self.X == x && self.Y == y {
		self.HandleCollision(target)
		return true
	}
	return false
}

func (self *Friend) Act() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rint := r.Intn(10)
	switch rint {
	case 0:
		self.MoveRight()
	case 1:
		self.MoveLeft()
	case 2:
		self.MoveUp()
	case 3:
		self.MoveDown()
	}
}
