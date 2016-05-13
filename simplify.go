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

	removedFaces := make(map[*Face]bool)

	// simplify
	numFaces := len(input.Triangles)
	target := int(float64(numFaces) * factor)
	for numFaces > target {
		// pop best pair
		p := heap.Pop(&queue).(*Pair)

		// move A to best position
		p.A.Vector = p.Vector()

		// update quadric matrix for A
		p.A.Quadric = p.Quadric()

		// consolidate faces
		distinctFaces := make(map[*Face]bool)
		for _, f := range vertexFaces[p.A] {
			distinctFaces[f] = true
		}
		for _, f := range vertexFaces[p.B] {
			distinctFaces[f] = true
		}

		// update faces and prune degenerate faces
		vertexFaces[p.A] = nil
		delete(vertexFaces, p.B)
		for f := range distinctFaces {
			if f.V1 == p.B {
				f.V1 = p.A
			}
			if f.V2 == p.B {
				f.V2 = p.A
			}
			if f.V3 == p.B {
				f.V3 = p.A
			}
			if f.Degenerate() {
				// TODO: remove f from vertexFaces[f.V1|f.V2|f.V3]
				if _, ok := removedFaces[f]; !ok {
					numFaces--
					removedFaces[f] = true
				}
				continue
			}
			vertexFaces[p.A] = append(vertexFaces[p.A], f)
		}

		// consolidate pairs
		distinctPairs := make(map[*Pair]bool)
		for _, q := range vertexPairs[p.A] {
			distinctPairs[q] = true
		}
		for _, q := range vertexPairs[p.B] {
			distinctPairs[q] = true
		}

		// update pairs and prune current pair
		vertexPairs[p.A] = nil
		delete(vertexPairs, p.B)
		seenPairs := make(map[PairKey]bool)
		for q := range distinctPairs {
			if q == p {
				continue
			}
			if q.A == p.B {
				q.A = p.A
			}
			if q.B == p.B {
				q.B = p.A
			}
			queue.Fix(q)
			if q.A == q.B {
				queue.Remove(q)
				continue
			}
			key := MakePairKey(q.A, q.B)
			if _, ok := seenPairs[key]; ok {
				queue.Remove(q)
				continue
			}
			seenPairs[key] = true
			vertexPairs[p.A] = append(vertexPairs[p.A], q)
		}
	}

	// find distinct faces
	distinctFaces := make(map[*Face]bool)
	for _, faces := range vertexFaces {
		for _, f := range faces {
			if !f.Degenerate() { // TODO: why does this happen?
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
