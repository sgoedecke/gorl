package main

import (
	"github.com/nsf/termbox-go"
	"strconv"
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

	// initialize the world and do initial draw
	w := NewWorld(80, 40)         // width 80, height 40
	p := Player{2, 2, 64, 100, w} // 64 -> '@'
	w.Player = &p
	l := Log{}
	hp := HealthBar{&p}

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

	w.Draw()

	l.AddMessage("Hello worllkjhlkjhlkjhlkjhlkjhlkjhlkjhlkjhlkjhlkjhlkjhlkjhlkjhlkjhlkjhlkjhlkjhlkjhjhjhjhdQQQ"+strconv.Itoa(len(l.Messages)), termbox.ColorCyan)

	l.Draw()
	hp.Draw(41, 80) // world height + 1, world width
	_ = termbox.Flush()
}
