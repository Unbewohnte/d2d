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

type Line struct {
	Start image.Point
	End   image.Point
}

func NewLine(start image.Point, end image.Point) Line {
	return Line{
		Start: start,
		End:   end,
	}
}

func (l *Line) Draw(canvas draw.Image, color color.Color) {
	var (
		x0      = l.Start.X
		x1      = l.End.X
		y0      = l.Start.Y
		y1      = l.End.Y
		dx  int = x1 - x0
		sx  int
		dy  int = y1 - y0
		sy  int
		err int
		e2  int
	)

	// abs dx
	if dx < 0 {
		dx = -dx
	}

	// -abs dy
	if dy < 0 {
		dy = -dy
	}
	dy = -dy

	// error
	err = dx + dy

	// sx
	if x0 < x1 {
		sx = 1
	} else {
		sx = -1
	}

	//sy
	if y0 < y1 {
		sy = 1
	} else {
		sy = -1
	}

	for {
		canvas.Set(x0, y0, color)
		if x0 == x1 && y0 == y1 {
			break
		}
		e2 = 2 * err
		if e2 >= dy {
			if x0 == x1 {
				break
			}
			err = err + dy
			x0 = x0 + sx
		}
		if e2 <= dx {
			if y0 == y1 {
				break
			}
			err = err + dx
			y0 = y0 + sy
		}
	}
}
