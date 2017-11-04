package main

import (
	"github.com/nsf/termbox-go"
)

func main() {
	// initialize termbox
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetOutputMode(termbox.Output256)

	// initialize an event queue and poll eternally, sending events to a channel
	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	// initialize the world and do initial draw
	w := NewWorld(80, 40)
	p := Player{2, 2, 64, w} // 64 -> '@'
	draw(w, p)

	// set up key handlers
	for {
		event := <-eventQueue
		if event.Type == termbox.EventKey {
			switch {
			case event.Key == termbox.KeyArrowLeft:
				p.MoveLeft()
			case event.Key == termbox.KeyArrowRight:
				p.MoveRight()
			case event.Key == termbox.KeyArrowUp:
				p.MoveUp()
			case event.Key == termbox.KeyArrowDown:
				p.MoveDown()
			case event.Ch == 'q' || event.Key == termbox.KeyEsc || event.Key == termbox.KeyCtrlC || event.Key == termbox.KeyCtrlD:
				return
			}
		}
		draw(w, p)

	}

}

// draw the world to the termbox back buffer & flush buffer
func draw(w *World, p Player) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for _, tile := range w.Tiles {
		termbox.SetCell(tile.X, tile.Y, tile.img, tile.fg, tile.bg)
	}

	termbox.SetCell(p.X, p.Y, p.img, termbox.ColorRed, termbox.ColorBlack)
	_ = termbox.Flush()
}
