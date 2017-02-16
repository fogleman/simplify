package simplify

import "math"

type PairKey struct {
	A, B Vector
}

func MakePairKey(a, b *Vertex) PairKey {
	if b.Less(a.Vector) {
		a, b = b, a
	}
	return PairKey{a.Vector, b.Vector}
}

type Pair struct {
	A, B        *Vertex
	Index       int
	Removed     bool
	CachedError float64
}

func NewPair(a, b *Vertex) *Pair {
	if b.Less(a.Vector) {
		a, b = b, a
	}
	return &Pair{a, b, -1, false, -1}
}

func (p *Pair) Quadric() Matrix {
	return p.A.Quadric.Add(p.B.Quadric)
}

func (p *Pair) Vector() Vector {
	q := p.Quadric()
	if math.Abs(q.Determinant()) > 1e-3 {
		v := q.QuadricVector()
		if !math.IsNaN(v.X) && !math.IsNaN(v.Y) && !math.IsNaN(v.Z) {
			return v
		}
	}
	// cannot compute best vector with matrix
	// look for best vector along edge
	const n = 32
	a := p.A.Vector
	b := p.B.Vector
	d := b.Sub(a)
	bestE := -1.0
	bestV := Vector{}
	for i := 0; i <= n; i++ {
		t := float64(i) / n
		v := a.Add(d.MulScalar(t))
		e := q.QuadricError(v)
		if bestE < 0 || e < bestE {
			bestE = e
			bestV = v
		}
	}
	return bestV
}

func (p *Pair) Error() float64 {
	if p.CachedError < 0 {
		p.CachedError = p.Quadric().QuadricError(p.Vector())
	}
	return p.CachedError
}
