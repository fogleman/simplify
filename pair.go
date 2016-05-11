package simplify

type PairKey struct {
	A, B Vector
}

type Pair struct {
	A, B Vertex
	Key  PairKey
}

func MakePair(a, b Vertex) Pair {
	if b.Less(a.Vector) {
		a, b = b, a
	}
	key := PairKey{a.Vector, b.Vector}
	return Pair{a, b, key}
}
