package main

import (
	"fmt"
	"os"

	gc "github.com/rthornton128/goncurses"
)

var (
	quit = make(chan struct{})
)

func main() {
	os.Exit(run(os.Args))
}

func run(args []string) int {
	if len(args) != 2 {
		fmt.Fprint(os.Stderr, "args is not 2")
		return 1
	}

	buf, err := initBuffer(args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	v, err := initView(buf)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	defer gc.End()

	if err := v.initScreen(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	if err := v.render(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	go func() {
		for {
			ch := v.win.GetChar()
			switch ch {
			case gc.KEY_UP, 'k':
				v.scrollUP()
			case gc.KEY_DOWN, '\n', 'j':
				v.scrollDOWN()
			case 'q':
				close(quit)
			case 0:
				continue
			default:
			}
		}
	}()

	for q := range quit {
		_ = q
		break
	}

	return 0
}
