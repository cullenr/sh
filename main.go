// Console version of a classic board game!

package main

import (
	"github.com/nsf/termbox-go"
	"time"
	"math/rand"
)

func randomLevel(w int, h int) [][]Tile {
	var out = make([][]Tile, h)
	for y := 0; y < h; y++ {
		out[y] = make([]Tile, w)
		for x := 0; x < w; x++ {
			if t := (x + y) % 2; t == 1 {
				out[y][x].Type = TileWall
			} else {
				out[y][x].Type = TileFloor
			}
		}
	}
	return out
}


func main() {
    debugDrawRects([]Rect64{{x: 10, y: 10, w: 100, h:100}, {x:110, y:110, w: 20, h:10}})
}

func main__() {
	rand.Seed(time.Now().UnixNano())

	board := randomLevel(32, 32)
	game := Game{board: board}
	// PollEvent is blocking so lets put it in to a goroutine
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

    render(&game.board, Point{x: 0, y: 0})

	for {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch {
				case ev.Key == termbox.KeyArrowLeft:
				case ev.Key == termbox.KeyArrowRight:
				case ev.Key == termbox.KeyArrowUp:
				case ev.Key == termbox.KeyArrowDown:
				case ev.Key == termbox.KeySpace:
				case ev.Ch == 'q' || ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyCtrlD:
					return
				}
			}
			// we can add network messages into this select too.
		default:
            render(&game.board, Point{x: 4, y: 4})
			// lets see what 60fps looks like
			time.Sleep(13 * time.Millisecond)
		}
	}
}
