package main

// Player type: an entity that does not try to attack the player
// Implements the DynamicEntity interface, uses the Entity mixin. Overrides
// collision-checking methods but borrows `Draw` and `Act` from Entity

type Player struct {
	Entity
}

// e is trying to move into the player!
func (self Player) HandleCollision(e DynamicEntity) {
	self.Log().AddMessage("Something hit you!", self.Color)
}

// e is trying to move into the player!
func (self Player) CollideWith(e DynamicEntity) {
	self.Log().AddMessage("You bumped into something!", self.Color)
}


func (self Player) Act() {
}
