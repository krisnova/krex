package trans

// goncurses - ncurses library for Go.
// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/* This example show a basic menu similar to that found in the ncurses
 * examples from TLDP */

import (
	. "github.com/rthornton128/goncurses"
)

const (
	DefaultHeight = 10
	DefaultWidth  = 30
)

type TransWindow struct {
	height int
	width  int
	my     int
	mx     int
	window *Window
	stdscr *Window
}

func GetNewWindow(height, width int) (*TransWindow, error) {

	stdscr, err := Init()
	if err != nil {
		return nil, err
	}
	my, mx := stdscr.MaxYX()
	y, x := 2, (mx/2)-(width/2)

	win, _ := NewWindow(height, width, y, x)
	win.Keypad(true)
	return &TransWindow{
		width:  width,
		height: height,
		window: win,
		my:     my,
		mx:     mx,
	}, nil
}

func (tw *TransWindow) StartScreen(msg string) error {
	stdscr, err := Init()
	if err != nil {
		return err
	}
	Raw(true)
	Echo(false)
	Cursor(0)
	stdscr.Clear()
	stdscr.Keypad(true)
	defer End()
	stdscr.Print(msg)
	stdscr.Refresh()
	tw.stdscr = stdscr
	return nil
}

func (tw *TransWindow) Prompt(items []string) {
	var active int
	printmenu(tw.window, items, active)

	for {
		ch := tw.stdscr.GetChar()
		switch Key(ch) {
		case 'q':
			return
		case KEY_UP:
			if active == 0 {
				active = len(items) - 1
			} else {
				active -= 1
			}
		case KEY_DOWN:
			if active == len(items)-1 {
				active = 0
			} else {
				active += 1
			}
		case KEY_RETURN, KEY_ENTER, Key('\r'):
			tw.stdscr.MovePrintf(tw.my-2, 0, "Choice #%d: %s selected",
				active,
				items[active])
			tw.stdscr.ClearToEOL()
			tw.stdscr.Refresh()
		default:
			tw.stdscr.MovePrintf(tw.my-2, 0, "Character pressed = %3d/%c",
				ch, ch)
			tw.stdscr.ClearToEOL()
			tw.stdscr.Refresh()
		}

		printmenu(tw.window, items, active)
	}
}

func printmenu(w *Window, menu []string, active int) {
	y, x := 2, 2
	w.Box(0, 0)
	for i, s := range menu {
		if i == active {
			w.AttrOn(A_REVERSE)
			w.MovePrint(y+i, x, s)
			w.AttrOff(A_REVERSE)
		} else {
			w.MovePrint(y+i, x, s)
		}
	}
	w.Refresh()
}
