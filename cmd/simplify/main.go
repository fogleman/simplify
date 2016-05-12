package main

import (
	"fmt"
	"os"

	"github.com/fogleman/simplify"
)

func main() {
	path := os.Args[1]
	fmt.Println(path)
	mesh, err := simplify.LoadBinarySTL(path)
	if err != nil {
		panic(err)
	}
	mesh = mesh.Simplify()
	mesh.SaveBinarySTL("out.stl")
}
