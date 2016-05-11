package simplify

type PairKey struct {
	A, B Vector
}

type Pair struct {
	A, B  Vertex
	Score float64
	Index int
}

func NewPair(a, b Vertex) *Pair {
	if b.Less(a.Vector) {
		a, b = b, a
	}
	q := a.Quadric.Add(b.Quadric)
	v := q.QuadricVector()
	s := q.QuadricError(v)
	return &Pair{a, b, s, -1}
}

func (p *Pair) Key() PairKey {
	return PairKey{p.A.Vector, p.B.Vector}
}
