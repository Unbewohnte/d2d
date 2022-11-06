package main

import (
	"image"
	"image/color"
	"unbewohnte/d2d"
	"unbewohnte/d2d/shapes"
)

const (
	width  int = 600
	height int = 600
)

func main() {
	var canvas = d2d.NewCanvas(uint16(width), uint16(height))
	canvas.FillWhole(color.RGBA{237, 223, 123, 255})

	canvas.DrawFilledRectangle(
		shapes.NewRectangle(
			image.Pt(0+width/10, 0),
			image.Pt(width/10+10, height),
		),
		color.RGBA{255, 190, 11, 255},
	)

	canvas.DrawFilledRectangle(
		shapes.NewRectangle(
			image.Pt(width-width/10, 0),
			image.Pt(width-width/10+10, height),
		),
		color.RGBA{255, 190, 11, 255},
	)

	var j uint8 = 0
	j = 0
	for i := width / 2; i > 1; i-- {
		canvas.DrawFilledCircle(
			shapes.NewCircle(
				image.Pt(width/2, height/2),
				i,
			),
			color.RGBA{131, 56 + j, 236, 255},
		)
		j++
	}

	for i := width / 3; i > 1; i-- {
		canvas.DrawFilledCircle(
			shapes.NewCircle(
				image.Pt(width/2, height/2),
				i,
			),
			color.RGBA{255, 190, 11 + j, 255},
		)
		j++
	}

	j = 0
	for i := width / 5; i > 1; i-- {
		canvas.DrawFilledCircle(
			shapes.NewCircle(
				image.Pt(width/2, height/2),
				i,
			),
			color.RGBA{251, 86 + j, 7, 255},
		)
		j++
	}

	err := canvas.SaveAsPNG("example_image1.png")
	if err != nil {
		panic(err)
	}
}
