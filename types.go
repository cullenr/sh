package main

type TileType int

const (
	TileOther TileType = iota
	TileWall
	TileFloor
)

type Actor interface {
}

type Point struct {
	x int
	y int
}

type Rect64 struct {
    // centre of the shape along x
    x float64
    // centre of the shape along y
    y float64
    w float64
    h float64
}

type Tile struct {
	Type  TileType
	actor *Actor
}

type Game struct {
	board [][]Tile
}

func (r *Rect64) left() float64 {
    return r.x - r.w * 0.5
}

func (r *Rect64) right() float64 {
    return r.x + r.w * 0.5
}

func (r *Rect64) top() float64 {
    return r.y - r.h * 0.5
}

func (r *Rect64) bottom() float64 {
    return r.y + r.h * 0.5
}
