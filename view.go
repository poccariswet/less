package main

import (
	gc "github.com/rthornton128/goncurses"
)

type View struct {
	buf        *Buffer
	win        *gc.Window
	height     int
	now_up_y   int
	now_down_y int
	max_y      int
}

func initView(buf *Buffer) (*View, error) {
	_, err := gc.Init()
	if err != nil {
		return nil, err
	}

	return &View{
		buf:        buf,
		now_up_y:   0,
		now_down_y: 0,
		max_y:      len(buf.data),
	}, nil
}

func (v *View) initScreen() error {
	stdscr := gc.StdScr()
	y, x := stdscr.MaxYX()

	win, err := gc.NewWindow(y, x, 0, 0)
	if err != nil {
		return err
	}
	v.win = win

	gc.Raw(true)
	gc.Cursor(0)
	gc.Echo(false)
	v.win.ScrollOk(true)
	v.win.Keypad(true)

	return nil
}

func (v *View) scrollUP() {
	if v.now_up_y >= 0 {
		v.win.Scroll(-1)
		v.win.MovePrint(0, 0, v.buf.data[v.now_up_y])
		v.win.Refresh()
		v.now_up_y--
		v.now_down_y--
	}
}

func (v *View) scrollDOWN() {
	if v.now_down_y < v.max_y {
		v.win.Scroll(1)
		v.win.MovePrintln(v.height-1, 0, v.buf.data[v.now_down_y])
		v.win.Refresh()
		v.now_down_y++
		v.now_up_y++
	}
}

func (v *View) render() error {
	y, _ := v.win.MaxYX()
	if y > len(v.buf.data) {
		y = len(v.buf.data)
	}
	y -= 1
	v.now_down_y = y
	v.height = y

	for i := 0; i < y; i++ {
		v.win.Println(v.buf.data[i])
		v.win.Refresh()
	}

	return nil
}
