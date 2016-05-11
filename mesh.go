package simplify

type Mesh struct {
	Triangles []*Triangle
}

func NewMesh(triangles []*Triangle) *Mesh {
	mesh := &Mesh{triangles}
	mesh.Initialize()
	return mesh
}

func (mesh *Mesh) Initialize() {
	quadrics := make(map[Vector]Matrix)
	triangles := make(map[Vector][]*Triangle)
	edges := make(map[Edge]bool)
	for _, t := range mesh.Triangles {
		m := t.Quadric()
		quadrics[t.V1] = quadrics[t.V1].Add(m)
		quadrics[t.V2] = quadrics[t.V2].Add(m)
		quadrics[t.V3] = quadrics[t.V3].Add(m)
		triangles[t.V1] = append(triangles[t.V1], t)
		triangles[t.V2] = append(triangles[t.V2], t)
		triangles[t.V3] = append(triangles[t.V3], t)
		edges[Edge{t.V1, t.V2}] = true
		edges[Edge{t.V2, t.V3}] = true
		edges[Edge{t.V3, t.V1}] = true
		edges[Edge{t.V2, t.V1}] = true
		edges[Edge{t.V3, t.V2}] = true
		edges[Edge{t.V1, t.V3}] = true
	}
	// for i := 0; i < 100; i++ {
	// 	for edge := range edges {
	// 		// fmt.Println(edge)
	// 		a, b := edge.A, edge.B
	// 		// c := a.Midpoint(b)
	// 		m := Matrix{}
	// 		for _, t := range triangles[a] {
	// 			m = m.Add(t.Quadric())
	// 		}
	// 		for _, t := range triangles[b] {
	// 			m = m.Add(t.Quadric())
	// 		}
	// 		// d := m.QuadricVector()
	// 		// fmt.Println(d)
	// 		// fmt.Println(m.QuadricError(a), m.QuadricError(b), m.QuadricError(c), m.QuadricError(d))
	// 		// fmt.Println()
	// 	}
	// }
}

func (m *Mesh) SaveBinarySTL(path string) error {
	return SaveBinarySTL(path, m)
}
