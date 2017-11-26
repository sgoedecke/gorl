package main

// Friend type: an entity that does not try to attack the player
// Implements the DynamicEntity interface, uses the Entity mixin. Overrides
// collision-checking methods but borrows `Draw` and `Act` from Entity

type Friend struct {
	Entity
}

func (self *Friend) HandleCollision(e DynamicEntity) {
	self.Log().AddMessage("Hello, friend!", self.Color)
}

func (self *Friend) Act() {
	self.RandomMove()
}
