package simplify

type Face struct {
	V1, V2, V3 *Vertex
	Removed    bool
}

func NewFace(v1, v2, v3 *Vertex) *Face {
	return &Face{v1, v2, v3, false}
}

func (f *Face) Area() float64 {
	a := f.V2.Sub(f.V1.Vector)
	b := f.V3.Sub(f.V1.Vector)
	return a.Cross(b).Length() / 2
}

func (f *Face) Degenerate() bool {
	return f.Area() <= 0
}
