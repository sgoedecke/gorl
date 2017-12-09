package main

import (
	"math/rand"
	"time"
)

// Enemy type: an entity that tries to attack the player
// Implements the DynamicEntity interface, uses the Entity mixin. Overrides
// collision-checking methods but borrows `Draw` and `Act` from Entity

type Enemy struct {
	Entity
}

func (self *Enemy) HandleCollision(e DynamicEntity) {
	self.Log().AddMessage("Oh no!", self.Color)
}

func (self *Enemy) CollideWith(e DynamicEntity) {
	self.Log().AddMessage("Have at you!", self.Color)
}

func (self *Enemy) Act() {
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
  rint := r.Intn(4)
  if rint == 0 {
    self.RandomMove()
  } else {
    self.ChasePlayer()
  }
}
