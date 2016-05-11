package simplify

import "fmt"

type Mesh struct {
	Triangles []*Triangle
}

func NewMesh(triangles []*Triangle) *Mesh {
	mesh := &Mesh{triangles}
	mesh.Initialize()
	return mesh
}

func (mesh *Mesh) Initialize() {
	// find distinct vertexes & associated triangles
	triangles := make(map[Vector][]*Triangle)
	for _, t := range mesh.Triangles {
		triangles[t.V1] = append(triangles[t.V1], t)
		triangles[t.V2] = append(triangles[t.V2], t)
		triangles[t.V3] = append(triangles[t.V3], t)
	}

	// make vertexes
	vertexes := make(map[Vector]Vertex)
	for v, t := range triangles {
		vertexes[v] = MakeVertex(v, t)
	}

	// find candidate pairs
	pairs := make(map[PairKey]Pair)
	for _, t := range mesh.Triangles {
		var pair Pair
		v1 := vertexes[t.V1]
		v2 := vertexes[t.V2]
		v3 := vertexes[t.V3]
		pair = MakePair(v1, v2)
		pairs[pair.Key] = pair
		pair = MakePair(v2, v3)
		pairs[pair.Key] = pair
		pair = MakePair(v3, v1)
		pairs[pair.Key] = pair
	}

	// TODO: pairs within a threshold distance

	fmt.Println(len(mesh.Triangles))
	fmt.Println(len(vertexes))
	fmt.Println(len(pairs))
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
