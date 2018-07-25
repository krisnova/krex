package trans

/*
 * The window file will ultimately be the glue between the rest
 * of krex and the user's terminal
 *
 * TODO @kris-nova Can we please make this beautiful and pretty and wonderful
 *
 */

import (
	"fmt"

	. "github.com/rthornton128/goncurses"
)

const (
	DefaultHeight = 30
	DefaultWidth  = 50
)

type TransWindow struct {
	height int
	width  int
	my     int
	mx     int

	// The menu window
	window *Window

	// The full terminal
	stdscr *Window
}

func GetNewWindow(height, width int) (*TransWindow, error) {
	stdscr, err := Init()
	if err != nil {
		return nil, err
	}
	my, mx := stdscr.Maxyx()

	// Offset of the inner window
	y, x := 2, (mx/2)-(width/2)

	fmt.Println(height, width, y, x)

	// 30 200 2 25
	win, _ := NewWindow(height, width, y, x)
	win.Keypad(true)

	Raw(true)
	Echo(false)
	Cursor(0)

	//stdscr.Clear()
	stdscr.Keypad(true)
	defer End()
	//stdscr.Print(msg)
	//stdscr.Refresh()

	return &TransWindow{
		width:  width,
		height: height,
		window: &win,
		stdscr: &stdscr,
		my:     my,
		mx:     mx,
	}, nil
}

func (tw *TransWindow) StartScreen(msg string) error {

	return nil
}

func (tw *TransWindow) Prompt(title string, items []string) string {

	// Init the prompt
	defer End()
	var active int

	// Clear the window
	//tw.window.Clear()
	tw.window.Refresh()

	// Clear the main screen

	//tw.stdscr.Clear()
	tw.stdscr.Refresh()
	tw.stdscr.Print(title)

	// Draw the inital window
	draw(tw.window, items, active)

	// Event loop
	for {
		ch := tw.stdscr.GetChar()
		switch Key(ch) {
		case 'q':
			//tw.stdscr.Clear()
			return ""
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
			tw.stdscr.MovePrint(tw.my-2, 0, "Choice #%d: %s selected",
				active,
				items[active])
			tw.stdscr.Refresh()
			tw.stdscr.Clear()
			return items[active]
		default:
			// Todo
			tw.stdscr.MovePrint(tw.my-2, 0, "Character pressed = %3d/%c",
				ch, ch)
			tw.stdscr.ClearToEOL()
			tw.stdscr.Refresh()
		}
		draw(tw.window, items, active)
	}
}

func (tw *TransWindow) End() {
	End()
}

func draw(w *Window, menu []string, active int) {
	y, x := 2, 2
	w.Box(0, 0)
	w.Background()
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
