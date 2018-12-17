package main

import(
    "github.com/fogleman/gg"
    "math/rand"
)

func debugDraw(rect Rect64) {
    dc := gg.NewContext(1000, 1000)
    dc.Scale(4, 4)
    dc.DrawRectangle(rect.x, rect.y, rect.w, rect.h)
    dc.SetRGB(rand.Float64(), rand.Float64(), rand.Float64())
    dc.Fill()
    dc.SavePNG("out.png")
}


func debugDrawRects(rects []Rect64) {
    dc := gg.NewContext(1000, 1000)
    dc.Scale(4, 4)
    for _, rect := range rects {
        dc.DrawRectangle(rect.x, rect.y, rect.w, rect.h)
        dc.SetRGB(rand.Float64(), rand.Float64(), rand.Float64())
        dc.Fill()
    }
    dc.SavePNG("out.png")
}
