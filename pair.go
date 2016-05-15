package simplify

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
	return p.Quadric().QuadricVector()
}

func (p *Pair) Error() float64 {
	if p.CachedError < 0 {
		q := p.Quadric()
		p.CachedError = q.QuadricError(q.QuadricVector())
	}
	return p.CachedError
}
