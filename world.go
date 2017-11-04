package main

import (
	"github.com/nsf/termbox-go"
)

type World struct {
	Tiles []Tile
}

type Tile struct {
	X        int
	Y        int
	img      rune
	fg       termbox.Attribute
	bg       termbox.Attribute
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
			// 184 is dark gray, 236 is '.'
			w.Tiles = append(w.Tiles, Tile{x, y, 184, 236, termbox.ColorDefault, true})
		}
	}

	// populate walls
	for y := 0; y < height; y++ {
		w.Tiles = append(w.Tiles, Tile{0, y, 35, termbox.ColorDefault, termbox.ColorDefault, false})         // 35 -> '#'
		w.Tiles = append(w.Tiles, Tile{width - 1, y, 35, termbox.ColorDefault, termbox.ColorDefault, false}) // '#'
	}

	for x := 0; x < width; x++ {
		w.Tiles = append(w.Tiles, Tile{x, 0, 35, termbox.ColorDefault, termbox.ColorDefault, false})          // '#'
		w.Tiles = append(w.Tiles, Tile{x, height - 1, 35, termbox.ColorDefault, termbox.ColorDefault, false}) // '#'
	}

	return &w
}
