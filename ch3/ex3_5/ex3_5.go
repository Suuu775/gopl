package ex35

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
	)
	pngfile, err := os.Create("mandelbrot.png")
	if err != nil {
		fmt.Println("file create fail")
		return
	}
	defer pngfile.Close()
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrotHelper(z))
		}
	}

	png.Encode(pngfile, img)
}

func mandelbrotHelper(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128

	for i := uint8(0); i < iterations; i++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.YCbCr{Y: i, Cb: 255 - contrast*i}
		}
	}
	return color.Black
}
