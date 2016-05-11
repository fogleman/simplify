package simplify

type Vertex struct {
	Vector
	Quadric   Matrix
	Triangles []*Triangle
}

func MakeVertex(vector Vector, triangles []*Triangle) Vertex {
	quadric := Matrix{}
	for _, t := range triangles {
		quadric = quadric.Add(t.Quadric())
	}
	return Vertex{vector, quadric, triangles}
}
