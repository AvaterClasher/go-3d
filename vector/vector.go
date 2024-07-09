package vector

import (
	"math"
)

type Vector struct {
	X, Y, Z float64
}

func Zero() Vector {
	return Vector{0, 0, 0}
}

func Unit() Vector {
	return Vector{1, 1, 1}
}

func UnitX() Vector {
	return Vector{1, 0, 0}
}

func UnitY() Vector {
	return Vector{0, 1, 0}
}

func UnitZ() Vector {
	return Vector{0, 0, 1}
}

func NewVector(x, y, z float64) Vector {
	return Vector{x, y, z}
}

func (v1 Vector) Add(v2 Vector) Vector {
	return Vector{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

func (v1 Vector) Sub(v2 Vector) Vector {
	return Vector{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z}
}

func (v1 Vector) Mul(v2 Vector) Vector {
	return Vector{v1.X * v2.X, v1.Y * v2.Y, v1.Z * v2.Z}
}

func (v1 Vector) Div(v2 Vector) Vector {
	return Vector{v1.X / v2.X, v1.Y / v2.Y, v1.Z / v2.Z}
}

func (v1 Vector) AddScalar(f float64) Vector {
	return Vector{v1.X + f, v1.Y + f, v1.Z + f}
}

func (v1 Vector) SubScalar(f float64) Vector {
	return Vector{v1.X - f, v1.Y - f, v1.Z - f}
}

func (v1 Vector) MulScalar(f float64) Vector {
	return Vector{v1.X * f, v1.Y * f, v1.Z * f}
}

func (v1 Vector) DivScalar(f float64) Vector {
	return Vector{v1.X / f, v1.Y / f, v1.Z / f}
}

func (v1 Vector) Dot(v2 Vector) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func (v1 Vector) Cross(v2 Vector) Vector {
	return Vector{
		v1.Y*v2.Z - v1.Z*v2.Y,
		v1.Z*v2.X - v1.X*v2.Z,
		v1.X*v2.Y - v1.Y*v2.X,
	}
}

func (v1 Vector) Length() float64 {
	return math.Sqrt(v1.X*v1.X + v1.Y*v1.Y + v1.Z*v1.Z)
}

func (v1 Vector) LengthSq() float64 {
	return v1.X*v1.X + v1.Y*v1.Y + v1.Z*v1.Z
}

func (v1 Vector) Normalize() Vector {
	return Vector{v1.X / v1.Length(), v1.Y / v1.Length(), v1.Z / v1.Length()}
}

func (v1 Vector) Negate() Vector {
	return Vector{-v1.X, -v1.Y, -v1.Z}
}