package go3d

import . "github.com/AvaterClasher/go-3d/vector"

type Mesh struct {
	Vertices []Vector
	Faces    []Face
}

func NewMesh() Mesh {
	return Mesh{}
}
