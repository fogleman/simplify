package simplify

type Triangle struct {
	V1, V2, V3 Vector
}

func NewTriangle(v1, v2, v3 Vector) *Triangle {
	return &Triangle{v1, v2, v3}
}

func (t *Triangle) Quadric() Matrix {
	n := t.Normal()
	x, y, z := t.V1.X, t.V1.Y, t.V1.Z
	a, b, c := n.X, n.Y, n.Z
	d := -a*x - b*y - c*z
	return Matrix{
		a * a, a * b, a * c, a * d,
		a * b, b * b, b * c, b * d,
		a * c, b * c, c * c, c * d,
		a * d, b * d, c * d, d * d,
	}
}

func (t *Triangle) Normal() Vector {
	e1 := t.V2.Sub(t.V1)
	e2 := t.V3.Sub(t.V1)
	return e1.Cross(e2).Normalize()
}
