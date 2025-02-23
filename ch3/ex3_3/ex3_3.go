package ex33

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func Surface() string {
	s := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)

			// Calculate average z value for color interpolation
			avgZ := (az + bz + cz + dz) / 4.0
			color := colorForZ(avgZ)

			s += fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	s += fmt.Sprintln("</svg>")
	return s
}

func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	if r == 0 {
		return 0 // Avoid division by zero
	}
	return math.Sin(r) / r
}

func colorForZ(z float64) string {
	// Normalize z value to [0, 1] range
	normalizedZ := (z + 1) / 2 // Assuming z ranges from -1 to 1
	if normalizedZ < 0 {
		normalizedZ = 0
	} else if normalizedZ > 1 {
		normalizedZ = 1
	}

	// Interpolate between blue (#0000ff) and red (#ff0000)
	blue := int((1 - normalizedZ) * 255)
	red := int(normalizedZ * 255)
	return fmt.Sprintf("#%02x%02x%02x", red, 0, blue)
}
