package main

import (
	"github.com/twada/golang-study/ch06/geometry"
	"github.com/twada/golang-study/ch06/coloredpoint"
	"fmt"
	"image/color"
)

func main() {
	{
		var cp coloredpoint.ColoredPoint
		cp.X = 1
		fmt.Println(cp.Point.X) // 1
		cp.Point.Y = 2
		fmt.Println(cp.Y) // 2
	}

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	{
		var p = coloredpoint.ColoredPoint{geometry.Point{1, 1}, red}
		var q = coloredpoint.ColoredPoint{geometry.Point{5, 4}, blue}
		fmt.Println(p.Distance(q.Point)) // 5
		p.ScaleBy(2)
		q.ScaleBy(2)
		fmt.Println(p.Distance(q.Point)) // 10
	}

	{
		var p = coloredpoint.ColoredPointP{&geometry.Point{1, 1}, red}
		var q = coloredpoint.ColoredPointP{&geometry.Point{5, 4}, blue}
		fmt.Println(p.Distance(*q.Point)) // 5
		q.Point = p.Point
		p.ScaleBy(2)
		fmt.Println(*p.Point, *q.Point) // {2 2} {2 2}
	}
}
