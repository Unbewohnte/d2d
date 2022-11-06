/*
The MIT License (MIT)

Copyright © 2022 Kasyanov Nikolay Alexeyevich (Unbewohnte)

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package shapes

import (
	"image"
	"image/color"
	"image/draw"
)

type Circle struct {
	Center image.Point
	Radius int
}

func NewCircle(center image.Point, radius int) Circle {
	if radius < 0 {
		radius = -radius
	}

	return Circle{
		Center: center,
		Radius: radius,
	}
}

func (c *Circle) Draw(canvas draw.Image, color color.Color) {
	var (
		x   int = c.Radius - 1
		y   int = 0
		dx  int = 1
		dy  int = 1
		err int = dx - (c.Radius * 2)
	)

	for x > y {
		canvas.Set(c.Center.X+x, c.Center.Y+y, color)
		canvas.Set(c.Center.X+y, c.Center.Y+x, color)
		canvas.Set(c.Center.X-y, c.Center.Y+x, color)
		canvas.Set(c.Center.X-x, c.Center.Y+y, color)
		canvas.Set(c.Center.X-x, c.Center.Y-y, color)
		canvas.Set(c.Center.X-y, c.Center.Y-x, color)
		canvas.Set(c.Center.X+y, c.Center.Y-x, color)
		canvas.Set(c.Center.X+x, c.Center.Y-y, color)

		if err <= 0 {
			y++
			err += dy
			dy += 2
		}
		if err > 0 {
			x--
			dx += 2
			err += dx - (c.Radius * 2)
		}
	}
}

func (c *Circle) DrawFilled(canvas draw.Image, color color.Color) {
	c.Draw(canvas, color)
	var (
		dx int
		dy int
	)

	for y := c.Center.Y - c.Radius; y < c.Center.Y+c.Radius; y++ {
		dy = y - c.Center.Y
		for x := c.Center.X - c.Radius; x < c.Center.X+c.Radius; x++ {
			dx = x - c.Center.X
			if (dx*dx + dy*dy) < (c.Radius * c.Radius) {
				canvas.Set(x, y, color)
			}
		}
	}
}
