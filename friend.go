package main

// Friend type: an entity that does not try to attack the player
// Implements the DynamicEntity interface, uses the Entity mixin. Overrides
// collision-checking methods but borrows `Draw` and `Act` from Entity

type Friend struct {
	Entity
}

func (self *Friend) HandleCollision(e *Entity, l *Log) {
	l.AddMessage("Hello, friend!", self.Color)
}

func (self *Friend) CheckCollision(target *Entity, x int, y int) bool {
	if self.X == x && self.Y == y {
		self.HandleCollision(target, self.screen.World.Log)
		return true
	}
	return false
}

func (self *Friend) Act() {
	self.Log().AddMessage("Moving!", self.Color)
	self.MoveRight()
}
