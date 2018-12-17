package main

import (
	"github.com/nsf/termbox-go"
)

const backgroundColor = termbox.ColorBlack

var colours = map[TileType]termbox.Attribute{
	TileWall:  termbox.ColorBlue,
	TileFloor: termbox.ColorGreen,
	TileOther: termbox.ColorMagenta,
}

func render(board *[][]Tile, offset Point) {
	termbox.Clear(backgroundColor, backgroundColor)
    for y, column := range *board {
		for x, tile := range column {
			bgColour := colours[tile.Type]
            // we draw two chars as we want to draw squares not rectangles
			termbox.SetCell((x + offset.x) * 2, (y + offset.y), ' ', bgColour, bgColour)
			termbox.SetCell((x + offset.x) * 2 + 1, (y + offset.y), ' ', bgColour, bgColour)
		}
	}
	termbox.Flush()
}
