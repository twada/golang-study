package coloredpoint

import (
	"github.com/twada/golang-study/ch06/geometry"
	"image/color"
)

type ColoredPoint struct {
	geometry.Point
	Color color.RGBA
}

type ColoredPointP struct {
	*geometry.Point
	Color color.RGBA
}
