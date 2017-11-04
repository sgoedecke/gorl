package main

type World struct {
	Tiles []Tile
}

type Tile struct {
	X        int
	Y        int
	img      rune
	passable bool
}

func (w *World) IsTileOccupied(x int, y int) bool {
	for _, tile := range w.Tiles {
		if tile.X == x && tile.Y == y && !tile.passable {
			return true
		}
	}
	return false
}

func NewWorld(width int, height int) *World {
	var w World
	// populate inner world with blank tiles
	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			w.Tiles = append(w.Tiles, Tile{x, y, 46, true}) // 46 -> '.'
		}
	}

	// populate walls
	for y := 0; y < height; y++ {
		w.Tiles = append(w.Tiles, Tile{0, y, 37, false})         // '%'
		w.Tiles = append(w.Tiles, Tile{width - 1, y, 37, false}) // '%'
	}

	for x := 0; x < width; x++ {
		w.Tiles = append(w.Tiles, Tile{x, 0, 37, false})          // '%'
		w.Tiles = append(w.Tiles, Tile{x, height - 1, 37, false}) // '%'
	}

	return &w
}
