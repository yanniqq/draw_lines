package main

import (
	"github.com/fogleman/gg"

	"math/rand"
)

func main() {
	const W = 1024
	const H = 1024
	drawContext := gg.NewContext(W, H)
	drawContext.SetRGB(0,0,0,)
	drawContext.Clear()
	for i := 0; i < 1000; i++ {
		x1 := rand.Float64() * W
		y1 := rand.Float64() * H
		x2 := rand.Float64() * W
		y2 := rand.Float64() * H
		r := rand.Float64()
		g := rand.Float64()
		b := rand.Float64()
		a := rand.Float64()*0.5 + 0.5
		w := rand.Float64()*4 + 1
		drawContext.SetRGBA(r, g, b, a)
		drawContext.SetLineWidth(w)
		drawContext.DrawLine(x1, y1, x2, y2)
		drawContext.Stroke()
	}
	drawContext.SavePNG("lines.png")
}