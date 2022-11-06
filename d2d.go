/*
The MIT License (MIT)

Copyright © 2022 Kasyanov Nikolay Alexeyevich (Unbewohnte)

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package d2d

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"unbewohnte/d2d/shapes"
)

// Image wrapper
type Canvas struct {
	innerImage draw.Image
}

// Creates new canvas with provided resolution
func NewCanvas(width uint16, height uint16) *Canvas {
	return &Canvas{
		innerImage: image.NewNRGBA(image.Rect(0, 0, int(width), int(height))),
	}
}

// Returns a pointer to the wrapped draw.Image
func (c *Canvas) InnerImage() *draw.Image {
	return &c.innerImage
}

// Fills the whole canvas with color
func (c *Canvas) FillWhole(color color.Color) {
	for y := 0; y < c.innerImage.Bounds().Dy(); y++ {
		for x := 0; x < c.innerImage.Bounds().Dx(); x++ {
			c.innerImage.Set(x, y, color)
		}
	}
}

// Swaps old neighboring colors with a new color
func (c *Canvas) FloodFill(pt image.Point, oldColor color.Color, newColor color.Color) {
	if c.innerImage.At(pt.X, pt.Y) == c.innerImage.ColorModel().Convert(oldColor) {
		c.innerImage.Set(pt.X, pt.Y, newColor)
		c.FloodFill(image.Pt(pt.X-1, pt.Y), oldColor, newColor)
		c.FloodFill(image.Pt(pt.X+1, pt.Y), oldColor, newColor)
		c.FloodFill(image.Pt(pt.X, pt.Y-1), oldColor, newColor)
		c.FloodFill(image.Pt(pt.X, pt.Y+1), oldColor, newColor)
	}
}

// Save image data as PNG file
func (c *Canvas) SaveAsPNG(filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}

	err = png.Encode(file, c.innerImage)
	if err != nil {
		return err
	}

	return nil
}

// Returns the bottom-right point of canvas
func (c *Canvas) Bounds() image.Point {
	return image.Pt(c.innerImage.Bounds().Dx(), c.innerImage.Bounds().Dy())
}

// Scales up image by magnitude
func (c *Canvas) ScaleUp(magnitude uint8) {
	var scaledCanvasImage *image.NRGBA = image.NewNRGBA(
		image.Rect(
			0,
			0,
			c.innerImage.Bounds().Dx()*int(magnitude),
			c.innerImage.Bounds().Dy()*int(magnitude),
		),
	)
	var basePixelColor color.Color
	for y0 := 0; y0 < c.innerImage.Bounds().Dy(); y0++ {
		for x0 := 0; x0 < c.innerImage.Bounds().Dx(); x0++ {
			basePixelColor = c.innerImage.At(x0, y0)

			for y := int(magnitude) * y0; y < int(magnitude)*(y0+1); y++ {
				for x := int(magnitude) * x0; x < int(magnitude)*(x0+1); x++ {
					scaledCanvasImage.Set(x, y, basePixelColor)
				}
			}
		}
	}

	c.innerImage = scaledCanvasImage
}

// Draws a point on canvas
func (c *Canvas) DrawPoint(pt image.Point, color color.Color) {
	c.innerImage.Set(pt.X, pt.Y, color)
}

// Draws a circle on canvas
func (c *Canvas) DrawCircle(circle shapes.Circle, color color.Color) {
	circle.Draw(c.innerImage, color)
}

// Draws a filled circle on canvas
func (c *Canvas) DrawFilledCircle(circle shapes.Circle, color color.Color) {
	circle.DrawFilled(c.innerImage, color)
}

// Draws a rectangle on canvas
func (c *Canvas) DrawRectangle(rectangle shapes.Rectangle, color color.Color) {
	rectangle.Draw(c.innerImage, color)
}

// Draws a filled rectangle on canvas
func (c *Canvas) DrawFilledRectangle(rectangle shapes.Rectangle, color color.Color) {
	rectangle.DrawFilled(c.innerImage, color)
}

// Draws a line on canvas
func (c *Canvas) DrawLine(line shapes.Line, color color.Color) {
	line.Draw(c.innerImage, color)
}

// Draws a triangle on canvas
func (c *Canvas) DrawTriangle(triangle shapes.Triangle, color color.Color) {
	triangle.Draw(c.innerImage, color)
}

// Draws a filled triangle on canvas
func (c *Canvas) DrawFilledTriangle(triangle shapes.Triangle, color color.Color) {
	triangle.DrawFilled(c.innerImage, color)
}

// Draws a grid on canvas
func (c *Canvas) Grid(border shapes.Rectangle, spacing uint16, lineWidth uint16, color color.Color) {
	if lineWidth == 0 {
		return
	}

	for x := border.UpperLeft.X; x < border.BottomRight.X; x += int(lineWidth + spacing) {
		c.DrawFilledRectangle(
			shapes.NewRectangle(
				image.Pt(x, border.UpperLeft.Y),
				image.Pt(x+int(lineWidth), border.BottomRight.Y),
			),
			color,
		)
	}

	for y := border.UpperLeft.Y; y < border.BottomRight.Y; y += int(lineWidth + spacing) {
		c.DrawFilledRectangle(
			shapes.NewRectangle(
				image.Pt(border.UpperLeft.X, y),
				image.Pt(border.BottomRight.X, y+int(lineWidth)),
			),
			color,
		)
	}
}

// Draws a canvas border
func (c *Canvas) Border(width uint16, color color.Color) {
	// upper part
	c.DrawFilledRectangle(
		shapes.NewRectangle(
			image.Pt(0, 0),
			image.Pt(c.innerImage.Bounds().Dx(), int(width)),
		),
		color,
	)

	// left part
	c.DrawFilledRectangle(
		shapes.NewRectangle(
			image.Pt(0, 0),
			image.Pt(int(width), c.innerImage.Bounds().Dy()),
		),
		color,
	)

	// right part
	c.DrawFilledRectangle(
		shapes.NewRectangle(
			image.Pt(c.innerImage.Bounds().Dx()-int(width), 0),
			image.Pt(c.innerImage.Bounds().Dx(), c.innerImage.Bounds().Dy()),
		),
		color,
	)

	// bottom part
	c.DrawFilledRectangle(
		shapes.NewRectangle(
			image.Pt(0, c.innerImage.Bounds().Dy()-int(width)),
			image.Pt(c.innerImage.Bounds().Dx(), c.innerImage.Bounds().Dy()),
		),
		color,
	)
}
