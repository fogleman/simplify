## 3D Mesh Simplification

Implementation of
[Surface Simplification Using Quadric Error Metrics, SIGGRAPH 97](http://cseweb.ucsd.edu/~ravir/190/2016/garland97.pdf),
written in Go.

![Bunny](http://i.imgur.com/Ttf5XzH.png)

<p align="center">270,000 faces vs. 2,700 faces (1%)</p>

---

### Install

    go get -u github.com/fogleman/simplify/cmd/simplify

### Command-Line Usage

    Usage: simplify [-f FACTOR] input.stl output.stl

    $ simplify -f 0.1 bunny.stl out.stl
    Loading bunny.stl
    Input mesh contains 270021 faces
    Simplifying to 10% of original...
    Output mesh contains 27001 faces
    Writing out.stl

### API Usage

```go
// Use LoadSTL (ASCII) or LoadBinarySTL
mesh, err := simplify.LoadBinarySTL(inputPath)
// handle err
mesh = mesh.Simplify(factor)
mesh.SaveBinarySTL(outputPath)
```

---

![Animated](http://i.imgur.com/LXke1ur.gif)

<p align="center">Iteratively simplifying by 50% until only 16 faces remain</p>
