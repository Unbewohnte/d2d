package main

import (
	"image"
	"image/color"
	"unbewohnte/d2d"
	"unbewohnte/d2d/shapes"
)

const (
	width  int = 1400
	height int = 800
)

func main() {
	var canvas = d2d.NewCanvas(uint16(width), uint16(height))
	canvas.FillWhole(color.RGBA{237, 223, 123, 255})

	canvas.Border(5, color.Black)
	canvas.DrawPoint(image.Pt(500, 700), color.Black)
	canvas.DrawFilledCircle(shapes.NewCircle(image.Pt(500, 500), 80), color.Black)
	canvas.DrawFilledRectangle(shapes.NewRectangle(image.Pt(20, 20), image.Pt(300, 300)), color.Black)
	canvas.DrawLine(shapes.NewLine(image.Pt(0, 0), image.Pt(1920, 1080)), color.Black)
	canvas.DrawFilledTriangle(
		shapes.NewTriangle(
			image.Pt(900, 20),
			image.Pt(1400, 400),
			image.Pt(500, 300),
		),
		color.Black,
	)

	err := canvas.SaveAsPNG("example_image2.png")
	if err != nil {
		panic(err)
	}
}
