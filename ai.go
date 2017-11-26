package main

import (
	"math/rand"
	"time"
)

func (self *Entity) RandomMove() {
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

func (self *Entity) ChasePlayer() {
  player := self.screen.World.Player
  if self.X < player.X {
    self.MoveRight()
    return
  }

  if self.X > player.X {
    self.MoveLeft()
    return
  }

  if self.Y > player.Y {
    self.MoveUp()
    return
  }

  if self.Y < player.Y {
    self.MoveDown()
    return
  }
}
