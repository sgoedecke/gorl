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

func (t *Tile) HandleCollision(e *Entity, l *Log) {
	l.AddMessage("You can't go there", termbox.ColorWhite)
}

// World type: has tiles. Can create itself with tiles as walls. Can check if there's an impassable tile at
// some coordinates

type World struct {
	Tiles    []Tile
	Player   *Entity
	Entities []*Entity
	Log      *Log
}

func (w *World) IsTileOccupied(e *Entity, x int, y int) bool {
	for _, tile := range w.Tiles {
		if tile.X == x && tile.Y == y && !tile.passable {
			tile.HandleCollision(e, w.Log)
			return true
		}
	}

	for _, entity := range w.Entities {
		if entity.X == x && entity.Y == y {
			entity.HandleCollision(e, w.Log)
			return true
		}
	}
	return false
}

func (w *World) Draw(x int, y int) {
	// draw all tiles
	for _, tile := range w.Tiles {
		termbox.SetCell(tile.X+x, tile.Y+y, tile.img, tile.fg, tile.bg)
	}
	// draw entities
	for _, entity := range w.Entities {
		termbox.SetCell(entity.X+x, entity.Y+y, entity.img, entity.Color, termbox.ColorBlack)
	}

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
