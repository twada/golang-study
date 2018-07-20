package geometry

import "math"

type Point struct{ X, Y float64 }

// 昔ながらの関数
func Distance(p, q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

// 同じだが、 Point 型のメソッドとして
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

// ポインタレシーバを持つメソッド
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
