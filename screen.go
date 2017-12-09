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

// check for impassable tiles, other entities, and the player.
// e is the entity trying to move into the tile (x,y)
func (s *Screen) IsTileOccupied(e DynamicEntity, x int, y int) bool {
	for _, tile := range s.Tiles {
		if tile.X == x && tile.Y == y && !tile.passable {
			return true
		}
	}

	// handle colliding with non-player entities (e might still be player)
	for i, _ := range s.Entities {
		entity := s.Entities[i] // get actual entity
		if entity.GetX() == x && entity.GetY() == y {
			player, ok := e.(Player) // if e is a player, trigger collision events
			if ok {
				player.CollideWith(entity)
				entity.HandleCollision(player)
			}
			return true
		}
	}

	// handle collisions with the player (e isn't player)
	player := s.World.Player
	if player.X == x && player.Y == y {
		e.CollideWith(player)
		player.HandleCollision(e)
		return true
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
