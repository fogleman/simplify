package simplify

import (
	"container/heap"
	"fmt"
)

type Mesh struct {
	Triangles []*Triangle
}

func NewMesh(triangles []*Triangle) *Mesh {
	return &Mesh{triangles}
}

func (m *Mesh) SaveBinarySTL(path string) error {
	return SaveBinarySTL(path, m)
}

func (mesh *Mesh) Simplify() *Mesh {
	// find distinct vertexes & associated triangles
	fmt.Println("find distinct vertexes & associated triangles")
	triangles := make(map[Vector][]*Triangle)
	for _, t := range mesh.Triangles {
		triangles[t.V1] = append(triangles[t.V1], t)
		triangles[t.V2] = append(triangles[t.V2], t)
		triangles[t.V3] = append(triangles[t.V3], t)
	}

	// make vertexes
	fmt.Println("make vertexes")
	vertexes := make(map[Vector]Vertex)
	for v, t := range triangles {
		vertexes[v] = MakeVertex(v, t)
	}

	// find candidate pairs
	fmt.Println("find candidate pairs")
	pairs := make(map[PairKey]*Pair)
	var queue PriorityQueue
	for _, t := range mesh.Triangles {
		v1 := vertexes[t.V1]
		v2 := vertexes[t.V2]
		v3 := vertexes[t.V3]
		var p *Pair
		p = NewPair(v1, v2)
		pairs[p.Key()] = p
		heap.Push(&queue, p)
		p = NewPair(v2, v3)
		pairs[p.Key()] = p
		heap.Push(&queue, p)
		p = NewPair(v3, v1)
		pairs[p.Key()] = p
		heap.Push(&queue, p)
	}

	// for len(queue) > 0 {
	// 	p := heap.Pop(&queue).(*Pair)
	// }

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

	return mesh
}
