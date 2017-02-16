package simplify

import "container/heap"

func Simplify(input *Mesh, factor float64) *Mesh {
	// find distinct vertices
	vectorVertex := make(map[Vector]*Vertex)
	for _, t := range input.Triangles {
		vectorVertex[t.V1] = NewVertex(t.V1)
		vectorVertex[t.V2] = NewVertex(t.V2)
		vectorVertex[t.V3] = NewVertex(t.V3)
	}

	// accumlate quadric matrices for each vertex based on its faces
	for _, t := range input.Triangles {
		q := t.Quadric()
		v1 := vectorVertex[t.V1]
		v2 := vectorVertex[t.V2]
		v3 := vectorVertex[t.V3]
		v1.Quadric = v1.Quadric.Add(q)
		v2.Quadric = v2.Quadric.Add(q)
		v3.Quadric = v3.Quadric.Add(q)
	}

	// create faces and map vertex => faces
	vertexFaces := make(map[*Vertex][]*Face)
	for _, t := range input.Triangles {
		v1 := vectorVertex[t.V1]
		v2 := vectorVertex[t.V2]
		v3 := vectorVertex[t.V3]
		f := NewFace(v1, v2, v3)
		vertexFaces[v1] = append(vertexFaces[v1], f)
		vertexFaces[v2] = append(vertexFaces[v2], f)
		vertexFaces[v3] = append(vertexFaces[v3], f)
	}

	// find distinct pairs
	// TODO: pair vertices within a threshold distance of each other
	pairs := make(map[PairKey]*Pair)
	for _, t := range input.Triangles {
		v1 := vectorVertex[t.V1]
		v2 := vectorVertex[t.V2]
		v3 := vectorVertex[t.V3]
		pairs[MakePairKey(v1, v2)] = NewPair(v1, v2)
		pairs[MakePairKey(v2, v3)] = NewPair(v2, v3)
		pairs[MakePairKey(v3, v1)] = NewPair(v3, v1)
	}

	// enqueue pairs and map vertex => pairs
	var queue PriorityQueue
	vertexPairs := make(map[*Vertex][]*Pair)
	for _, p := range pairs {
		heap.Push(&queue, p)
		vertexPairs[p.A] = append(vertexPairs[p.A], p)
		vertexPairs[p.B] = append(vertexPairs[p.B], p)
	}

	// simplify
	numFaces := len(input.Triangles)
	target := int(float64(numFaces) * factor)
	for numFaces > target {
		// pop best pair
		p := heap.Pop(&queue).(*Pair)

		if p.Removed {
			continue
		}
		p.Removed = true

		// get related faces
		distinctFaces := make(map[*Face]bool)
		for _, f := range vertexFaces[p.A] {
			if !f.Removed {
				distinctFaces[f] = true
			}
		}
		for _, f := range vertexFaces[p.B] {
			if !f.Removed {
				distinctFaces[f] = true
			}
		}

		// get related pairs
		distinctPairs := make(map[*Pair]bool)
		for _, q := range vertexPairs[p.A] {
			if !q.Removed {
				distinctPairs[q] = true
			}
		}
		for _, q := range vertexPairs[p.B] {
			if !q.Removed {
				distinctPairs[q] = true
			}
		}

		// create the new vertex
		v := &Vertex{p.Vector(), p.Quadric()}

		// update faces
		newFaces := make([]*Face, 0, len(distinctFaces))
		valid := true
		for f := range distinctFaces {
			v1, v2, v3 := f.V1, f.V2, f.V3
			if v1 == p.A || v1 == p.B {
				v1 = v
			}
			if v2 == p.A || v2 == p.B {
				v2 = v
			}
			if v3 == p.A || v3 == p.B {
				v3 = v
			}
			face := NewFace(v1, v2, v3)
			if face.Degenerate() {
				continue
			}
			if face.Normal().Dot(f.Normal()) < 1e-3 {
				valid = false
				break
			}
			newFaces = append(newFaces, face)
		}
		if !valid {
			continue
		}
		delete(vertexFaces, p.A)
		delete(vertexFaces, p.B)
		for f := range distinctFaces {
			f.Removed = true
			numFaces--
		}
		for _, f := range newFaces {
			numFaces++
			vertexFaces[f.V1] = append(vertexFaces[f.V1], f)
			vertexFaces[f.V2] = append(vertexFaces[f.V2], f)
			vertexFaces[f.V3] = append(vertexFaces[f.V3], f)
		}

		// update pairs and prune current pair
		delete(vertexPairs, p.A)
		delete(vertexPairs, p.B)
		seen := make(map[Vector]bool)
		for q := range distinctPairs {
			q.Removed = true
			heap.Remove(&queue, q.Index)
			a, b := q.A, q.B
			if a == p.A || a == p.B {
				a = v
			}
			if b == p.A || b == p.B {
				b = v
			}
			if b == v {
				// swap so that a == v
				a, b = b, a
			}
			if _, ok := seen[b.Vector]; ok {
				// only want distinct neighbors
				continue
			}
			seen[b.Vector] = true
			q = NewPair(a, b)
			heap.Push(&queue, q)
			vertexPairs[a] = append(vertexPairs[a], q)
			vertexPairs[b] = append(vertexPairs[b], q)
		}
	}

	// find distinct faces
	distinctFaces := make(map[*Face]bool)
	for _, faces := range vertexFaces {
		for _, f := range faces {
			if !f.Removed {
				distinctFaces[f] = true
			}
		}
	}

	// construct resulting mesh
	triangles := make([]*Triangle, len(distinctFaces))
	i := 0
	for f := range distinctFaces {
		triangles[i] = NewTriangle(f.V1.Vector, f.V2.Vector, f.V3.Vector)
		i++
	}
	return NewMesh(triangles)
}
