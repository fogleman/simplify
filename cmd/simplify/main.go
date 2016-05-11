package main

import (
	"fmt"
	"os"

	"github.com/fogleman/simplify"
)

func main() {
	path := os.Args[1]
	mesh, err := simplify.LoadBinarySTL(path)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(mesh.Triangles))
}
