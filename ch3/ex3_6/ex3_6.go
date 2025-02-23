package mandelbrot

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func Mandelbrot() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		iterations             = 200
		contrast               = 15
		supersample            = 2
	)

	pngfile, err := os.Create("mandelbrot.png")
	if err != nil {
		fmt.Println("file create fail")
		return
	}
	defer pngfile.Close()

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			var totalColor color.Color
			for i := 0; i < supersample; i++ {
				for j := 0; j < supersample; j++ {
					x := float64(px*supersample+i)/float64(width*supersample)*(xmax-xmin) + xmin
					y := float64(py*supersample+j)/float64(height*supersample)*(ymax-ymin) + ymin
					z := complex(x, y)
					color := mandelbrotHelper(z, iterations, contrast)
					if totalColor == nil {
						totalColor = color
					} else {
						totalColor = averageColors(totalColor, color)
					}
				}
			}
			img.Set(px, py, totalColor)
		}
	}

	png.Encode(pngfile, img)
}

func mandelbrotHelper(z complex128, iterations uint8, contrast uint8) color.Color {
	var v complex128
	for i := uint8(0); i < iterations; i++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}

func averageColors(c1, c2 color.Color) color.Color {
	r1, g1, b1, _ := c1.RGBA()
	r2, g2, b2, _ := c2.RGBA()

	r := (r1 + r2) / 2
	g := (g1 + g2) / 2
	b := (b1 + b2) / 2

	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
}
