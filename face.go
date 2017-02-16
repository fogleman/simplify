package simplify

type Face struct {
	V1, V2, V3 *Vertex
	Removed    bool
}

func NewFace(v1, v2, v3 *Vertex) *Face {
	return &Face{v1, v2, v3, false}
}

func (f *Face) Degenerate() bool {
	v1 := f.V1.Vector
	v2 := f.V2.Vector
	v3 := f.V3.Vector
	return v1 == v2 || v1 == v3 || v2 == v3
}

func (f *Face) Normal() Vector {
	e1 := f.V2.Sub(f.V1.Vector)
	e2 := f.V3.Sub(f.V1.Vector)
	return e1.Cross(e2).Normalize()
}
