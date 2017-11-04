package main

import (
	"github.com/nsf/termbox-go"
)

const LogWidth = 40
const WorldWidth = 80
const MaxHeight = 40

// Message type: an individual line in the log. Has a color. Text is stored as an array of strings
// (each which must be 40 chars or less) to make handling word wrap easy. Can draw itself if given a y.

type Message struct {
	Text  []string // array of length 40 strings to ensure line breaks work
	Color termbox.Attribute
}

func NewMessage(s string, color termbox.Attribute) *Message {
	var m Message
	m.Color = color

	if len(s) <= 40 {
		m.Text = append(m.Text, s)
	} else {
		slices := float32(len(s)) / 40.0 // number of slices
		// if slices is not a natural number, increment by 1 so the float rounding doesn't screw us
		if float32(int(slices)) < slices {
			slices = slices + 1
		}
		for i := 0; i < int(slices); i++ {
			left := 40 * i // we want to take slices of our string by 40 chars. set indices appropriately
			right := 40 * (i + 1)
			if right > len(s) { // make sure we don't slice a longer slice than our string
				right = len(s)
			}
			m.Text = append(m.Text, s[left:right])
		}
	}
	return &m
}

func (m *Message) Height() int {
	return len(m.Text)
}

func (m *Message) Draw(y int) {
	for _, str := range m.Text {
		x := int(WorldWidth + 1)
		runes := []rune(str)
		for _, r := range runes {
			termbox.SetCell(x, y, r, m.Color, termbox.ColorDefault)
			x++
		}
		y++
	}
}

// Log type: contains a lot of messages, knows how to add a new message to the bottom (deleting top ones
// for space). Knows how to draw itself.

type Log struct {
	Messages []Message
}

func (l *Log) Height() int {
	acc := 0
	for _, msg := range l.Messages {
		acc = acc + msg.Height()
	}
	return acc
}

func (l *Log) AddMessage(s string, color termbox.Attribute) {
	m := NewMessage(s, color)
	overflow := (l.Height() + m.Height()) - 40
	if overflow > 0 {
		l.Messages = l.Messages[overflow:] // tidy up
	}
	l.Messages = append(l.Messages, *m)
}

func (l *Log) Draw() {
	y := 0
	for _, msg := range l.Messages {
		msg.Draw(y)
		y = y + msg.Height()
	}
}
