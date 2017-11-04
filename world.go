package main

import (
	"github.com/nsf/termbox-go"
)

// Tile type: has coords, a rune, colors, and knows if you can step on it or not

type Tile struct {
	X        int
	Y        int
	img      rune
	fg       termbox.Attribute
	bg       termbox.Attribute
	passable bool
}

// World type: has tiles. Can create itself with tiles as walls. Can check if there's an impassable tile at
// some coordinates

type World struct {
	Tiles  []Tile
	Player *Player
}

func (w *World) IsTileOccupied(x int, y int) bool {
	for _, tile := range w.Tiles {
		if tile.X == x && tile.Y == y && !tile.passable {
			return true
		}
	}
	return false
}

func (w *World) Draw() {
	// draw all tiles
	for _, tile := range w.Tiles {
		termbox.SetCell(tile.X, tile.Y, tile.img, tile.fg, tile.bg)
	}
	// draw player
	p := w.Player
	termbox.SetCell(p.X, p.Y, p.img, termbox.ColorRed, termbox.ColorBlack)
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
