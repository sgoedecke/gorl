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
	termbox.SetOutputMode(termbox.Output256) // set 256-color mode

	// initialize an event queue and poll eternally, sending events to a channel
	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	// initialize the world, player and do initial draw
	w := NewWorld(80, 40)         // width 80, height 40
	p := Entity{2, 2, 64, 100, w} // 64 -> '@'
	w.Player = &p
	l := Log{}
	hp := HealthBar{&p}

	l.AddMessage("Welcome to Gorl!", termbox.ColorGreen)
	l.AddMessage("But be careful...", termbox.ColorRed)

	draw(w, &l, &hp)
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
		draw(w, &l, &hp)

	}

}

// clear buffer, draw everything to the termbox back buffer & flush buffer
// takes pointers to a world widget, a log widget, and a health bar widget
func draw(w *World, l *Log, hp *HealthBar) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	w.Draw(0, 0)   // draw world at 0,0
	l.Draw(81, 0)  // draw log to right of world
	hp.Draw(0, 41) // draw hp at the very bottom of the world
	_ = termbox.Flush()
}
