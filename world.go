package main

type World struct {
  ActiveScreen *Screen
	Player   *Entity
	Log      *Log
}

func NewWorld(width int, height int) *World {
	var w World

  w.ActiveScreen = NewScreen(width, height, &w)

	return &w
}

// delegate drawing world to the active screen
func (w *World) Draw(x int, y int) {
  w.ActiveScreen.Draw(x,y)
}
