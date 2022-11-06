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

type Rectangle struct {
	UpperLeft   image.Point
	BottomRight image.Point
}

func NewRectangle(upperLeft image.Point, bottomRight image.Point) Rectangle {
	return Rectangle{
		UpperLeft:   upperLeft,
		BottomRight: bottomRight,
	}
}

func (r *Rectangle) Draw(canvas draw.Image, color color.Color) {
	for y := r.UpperLeft.Y; y <= r.BottomRight.Y; y++ {
		canvas.Set(r.UpperLeft.X, y, color)
		canvas.Set(r.BottomRight.X, y, color)
	}

	for x := r.UpperLeft.X; x <= r.BottomRight.X; x++ {
		canvas.Set(x, r.UpperLeft.Y, color)
		canvas.Set(x, r.BottomRight.Y, color)
	}
}

func (r *Rectangle) DrawFilled(canvas draw.Image, color color.Color) {
	for y := 0; y < r.BottomRight.Y-r.UpperLeft.Y; y++ {
		for x := 0; x < r.BottomRight.X-r.UpperLeft.X; x++ {
			canvas.Set(x+r.UpperLeft.X, y+r.UpperLeft.Y, color)
		}
	}

}
