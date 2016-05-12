package simplify

type Mesh struct {
	Triangles []*Triangle
}

func NewMesh(triangles []*Triangle) *Mesh {
	return &Mesh{triangles}
}

func (m *Mesh) SaveBinarySTL(path string) error {
	return SaveBinarySTL(path, m)
}

func (m *Mesh) Simplify() *Mesh {
	return Simplify(m)
}
