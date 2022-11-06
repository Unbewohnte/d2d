# d2d - draw 2D
## the most minimal, probably most lacking, but simple 2d drawing Go package

![Example image](https://unbewohnte.su:3000/Unbewohnte/d2d/examples/example_image2.png)

## What can it do ?

Draw 

- Points
- Lines
- Rectangles (Filled and empty)
- Circles (Filled and empty)

of any color that satisfies basic stdlib's color.Color interface

## Installation

`go get unbewohnte.su:3000/Unbewohnte/d2d` - if it doesn't work, your best choice is to just download it manually or copy-paste the code into your project

## Usage

The main part of the package - `Canvas` struct that is essentially just a wrapper to `draw.Image`. Though the drawing can be done via direct work with shapes, canvas comes with additional quality of life drawing functions such as `Border`, `FillWhole`, `Grid`, `FloodFill`.  


### Example 

```
canvas := d2d.NewCanvas(1920, 1080)
canvas.FillWhole(color.RGBA{60, 180, 30, 255})
canvas.Border(5, color.Black)
canvas.DrawFilledCircle(shapes.NewCircle(image.Pt(500, 500), 80), color.Black)
canvas.DrawLine(shapes.NewLine(image.Pt(0, 0), image.Pt(1920, 1080)), color.Black)
canvas.DrawFilledTriangle(
    shapes.NewTriangle(
        image.Pt(900, 20),
        image.Pt(1400, 400),
        image.Pt(500, 300),
    ),
    color.Black,
)

err := canvas.SaveAsPNG("image.png")
if err != nil {
    panic(err)
}
```

## License

MIT