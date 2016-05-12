package simplify

type Vertex struct {
	Vector
	Quadric Matrix
}

func NewVertex(v Vector) *Vertex {
	return &Vertex{Vector: v}
}
