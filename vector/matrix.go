package vector

import "math"

type Matrix struct {
	M00, M01, M02, M03 float64
	M10, M11, M12, M13 float64
	M20, M21, M22, M23 float64
	M30, M31, M32, M33 float64
}

func Identity() Matrix {
	return Matrix{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func RotateX(angle float64) Matrix {
	s := math.Sin(angle)
	c := math.Cos(angle)
	return Matrix{
		1, 0, 0, 0,
		0, c, -s, 0,
		0, s, c, 0,
		0, 0, 0, 1,
	}
}

func RotateY(angle float64) Matrix {
	s := math.Sin(angle)
	c := math.Cos(angle)
	return Matrix{
		c, 0, s, 0,
		0, 1, 0, 0,
		-s, 0, c, 0,
		0, 0, 0, 1,
	}
}

func RotateZ(angle float64) Matrix {
	s := math.Sin(angle)
	c := math.Cos(angle)
	return Matrix{
		c, -s, 0, 0,
		s, c, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func RotateAxis(axis Vector, angle float64) Matrix {
	s := math.Sin(angle)
	c := math.Cos(angle)
	v := axis.Normalize()
	return Matrix{
		v.X*v.X*(1-c) + c,
		v.X*v.Y*(1-c) - v.Z*s,
		v.X*v.Z*(1-c) + v.Y*s,
		0,
		v.X*v.Y*(1-c) + v.Z*s,
		v.Y*v.Y*(1-c) + c,
		v.Y*v.Z*(1-c) - v.X*s,
		0,
		v.X*v.Z*(1-c) - v.Y*s, v.Y*v.Z*(1-c) + v.X*s, v.Z*v.Z*(1-c) + c, 0,
		0, 0, 0, 1,
	}
}

func Translate(v Vector) Matrix {
	return Matrix{
		1, 0, 0, v.X,
		0, 1, 0, v.Y,
		0, 0, 1, v.Z,
		0, 0, 0, 1,
	}
}

func Scale(v Vector) Matrix {
	return Matrix{
		v.X, 0, 0, 0,
		0, v.Y, 0, 0,
		0, 0, v.Z, 0,
		0, 0, 0, 1,
	}
}

func LookAt(from, to, up Vector) Matrix {
	z := from.Sub(to).Normalize()
	x := up.Cross(z).Normalize()
	y := z.Cross(x)
	return Matrix{
		x.X, x.Y, x.Z, -x.Dot(from),
		y.X, y.Y, y.Z, -y.Dot(from),
		z.X, z.Y, z.Z, -z.Dot(from),
		0, 0, 0, 1,
	}
}

