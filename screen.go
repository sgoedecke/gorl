package main

import (
	"github.com/nsf/termbox-go"
)

// Tile type: has coords, a rune, colors, and knows if you can step on it or not
// Aside from the passable field, this is purely for graphics

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

// Screen type: has tiles. Can create itself with tiles as walls. Can check if there's an impassable tile at
// some coordinates

type Screen struct {
	Tiles    []Tile
	Entities []DynamicEntity
	World    *World
}

func (s *Screen) Act() {
	for i, _ := range s.Entities {
		s.Entities[i].Act()
	}
}

func (s *Screen) IsTileOccupied(e *Entity, x int, y int) bool {
	for _, tile := range s.Tiles {
		if tile.X == x && tile.Y == y && !tile.passable {
			tile.HandleCollision(e, s.World.Log)
			return true
		}
	}

	for _, entity := range s.Entities {
		if entity.CheckCollision(e, x, y) {
			return true
		}
	}
	return false
}

func (s *Screen) Draw(x int, y int) {
	// draw all tiles
	for _, tile := range s.Tiles {
		termbox.SetCell(tile.X+x, tile.Y+y, tile.img, tile.fg, tile.bg)
	}
	// draw entities
	for _, entity := range s.Entities {
		entity.Draw(x, y)
	}

	// draw player
	player := s.World.Player
	termbox.SetCell(player.X+x, player.Y+y, player.img, player.Color, termbox.ColorBlack)

}

func NewScreen(width int, height int, w *World) *Screen {
	var s Screen
	s.World = w
	// populate inner world with blank tiles
	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			// 184 is dark gray, 236 is '.'
			s.Tiles = append(s.Tiles, Tile{x, y, 184, 236, termbox.ColorDefault, true})
		}
	}

	// populate walls
	for y := 0; y < height; y++ {
		s.Tiles = append(s.Tiles, Tile{0, y, 35, termbox.ColorDefault, termbox.ColorDefault, false})         // 35 -> '#'
		s.Tiles = append(s.Tiles, Tile{width - 1, y, 35, termbox.ColorDefault, termbox.ColorDefault, false}) // '#'
	}

	for x := 0; x < width; x++ {
		s.Tiles = append(s.Tiles, Tile{x, 0, 35, termbox.ColorDefault, termbox.ColorDefault, false})          // '#'
		s.Tiles = append(s.Tiles, Tile{x, height - 1, 35, termbox.ColorDefault, termbox.ColorDefault, false}) // '#'
	}

	return &s
}
