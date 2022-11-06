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
	"math"
)

type Triangle struct {
	p0 image.Point
	p1 image.Point
	p2 image.Point
}

func NewTriangle(p0 image.Point, p1 image.Point, p2 image.Point) Triangle {
	return Triangle{
		p0: p0,
		p1: p1,
		p2: p2,
	}
}

func (t *Triangle) Draw(canvas draw.Image, color color.Color) {
	line0 := NewLine(t.p0, t.p1)
	line0.Draw(canvas, color)

	line1 := NewLine(t.p1, t.p2)
	line1.Draw(canvas, color)

	line11 := NewLine(t.p2, t.p1)
	line11.Draw(canvas, color)

	line2 := NewLine(t.p2, t.p0)
	line2.Draw(canvas, color)
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func area(p0 image.Point, p1 image.Point, p2 image.Point) float32 {
	return float32(
		math.Abs(
			float64((p0.X*(p1.Y-p2.Y) + p1.X*(p2.Y-p0.Y) + p2.X*(p0.Y-p1.Y))) / 2.0,
		),
	)
}

func (t *Triangle) isPointInside(pt image.Point) bool {
	var (
		a  = area(t.p0, t.p1, t.p2)
		a1 = area(pt, t.p1, t.p2)
		a2 = area(t.p0, pt, t.p2)
		a3 = area(t.p0, t.p1, pt)
	)

	return (a == a1+a2+a3)
}

func (t *Triangle) DrawFilled(canvas draw.Image, color color.Color) {
	t.Draw(canvas, color)

	maxX := max(t.p0.X, max(t.p1.X, t.p2.X))
	maxY := max(t.p0.Y, max(t.p1.Y, t.p2.Y))
	minX := min(t.p0.X, min(t.p1.X, t.p2.X))
	minY := min(t.p0.Y, min(t.p1.Y, t.p2.Y))

	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			if t.isPointInside(image.Pt(x, y)) {
				canvas.Set(x, y, color)
			}
		}
	}
}
