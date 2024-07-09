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

func Orthographic(left, right, bottom, top, near, far float64) Matrix {
	return Matrix{
		2 / (right - left), 0, 0, -(right + left) / (right - left),
		0, 2 / (top - bottom), 0, -(top + bottom) / (top - bottom),
		0, 0, -2 / (far - near), -(far + near) / (far - near),
		0, 0, 0, 1,
	}
}

func Frustum(left, right, bottom, top, near, far float64) Matrix {
	return Matrix{
		(2 * near) / (right - left), 0, (right + left) / (right - left), 0,
		0, (2 * near) / (top - bottom), (top + bottom) / (top - bottom), 0,
		0, 0, -(far + near) / (far - near), (-2 * near * far) / (far - near),
		0, 0, -1, 0,
	}
}

func Perspective(fovy, aspect, near, far float64) Matrix {
	top := math.Tan(fovy*math.Pi/360) * near
	bottom := -top
	right := top * aspect
	left := bottom * aspect
	return Frustum(left, right, bottom, top, near, far)
}

func (m1 Matrix) Add(m2 Matrix) Matrix {
	return Matrix{
		m1.M00 + m2.M00, m1.M01 + m2.M01, m1.M02 + m2.M02, m1.M03 + m2.M03,
		m1.M10 + m2.M10, m1.M11 + m2.M11, m1.M12 + m2.M12, m1.M13 + m2.M13,
		m1.M20 + m2.M20, m1.M21 + m2.M21, m1.M22 + m2.M22, m1.M23 + m2.M23,
		m1.M30 + m2.M30, m1.M31 + m2.M31, m1.M32 + m2.M32, m1.M33 + m2.M33,
	}
}

func (m1 Matrix) Sub(m2 Matrix) Matrix {
	return Matrix{
		m1.M00 - m2.M00, m1.M01 - m2.M01, m1.M02 - m2.M02, m1.M03 - m2.M03,
		m1.M10 - m2.M10, m1.M11 - m2.M11, m1.M12 - m2.M12, m1.M13 - m2.M13,
		m1.M20 - m2.M20, m1.M21 - m2.M21, m1.M22 - m2.M22, m1.M23 - m2.M23,
		m1.M30 - m2.M30, m1.M31 - m2.M31, m1.M32 - m2.M32, m1.M33 - m2.M33,
	}
}

func (m1 Matrix) Mul(m2 Matrix) Matrix {
	return Matrix{
		m1.M00*m2.M00 + m1.M01*m2.M10 + m1.M02*m2.M20 + m1.M03*m2.M30,
		m1.M00*m2.M01 + m1.M01*m2.M11 + m1.M02*m2.M21 + m1.M03*m2.M31,
		m1.M00*m2.M02 + m1.M01*m2.M12 + m1.M02*m2.M22 + m1.M03*m2.M32,
		m1.M00*m2.M03 + m1.M01*m2.M13 + m1.M02*m2.M23 + m1.M03*m2.M33,
		m1.M10*m2.M00 + m1.M11*m2.M10 + m1.M12*m2.M20 + m1.M13*m2.M30,
		m1.M10*m2.M01 + m1.M11*m2.M11 + m1.M12*m2.M21 + m1.M13*m2.M31,
		m1.M10*m2.M02 + m1.M11*m2.M12 + m1.M12*m2.M22 + m1.M13*m2.M32,
		m1.M10*m2.M03 + m1.M11*m2.M13 + m1.M12*m2.M23 + m1.M13*m2.M33,
		m1.M20*m2.M00 + m1.M21*m2.M10 + m1.M22*m2.M20 + m1.M23*m2.M30,
		m1.M20*m2.M01 + m1.M21*m2.M11 + m1.M22*m2.M21 + m1.M23*m2.M31,
		m1.M20*m2.M02 + m1.M21*m2.M12 + m1.M22*m2.M22 + m1.M23*m2.M32,
		m1.M20*m2.M03 + m1.M21*m2.M13 + m1.M22*m2.M23 + m1.M23*m2.M33,
		m1.M30*m2.M00 + m1.M31*m2.M10 + m1.M32*m2.M20 + m1.M33*m2.M30,
		m1.M30*m2.M01 + m1.M31*m2.M11 + m1.M32*m2.M21 + m1.M33*m2.M31,
		m1.M30*m2.M02 + m1.M31*m2.M12 + m1.M32*m2.M22 + m1.M33*m2.M32,
		m1.M30*m2.M03 + m1.M31*m2.M13 + m1.M32*m2.M23 + m1.M33*m2.M33,
	}
}

func (m1 Matrix) MulScalar(f float64) Matrix {
	return Matrix{
		m1.M00 * f, m1.M01 * f, m1.M02 * f, m1.M03 * f,
		m1.M10 * f, m1.M11 * f, m1.M12 * f, m1.M13 * f,
		m1.M20 * f, m1.M21 * f, m1.M22 * f, m1.M23 * f,
		m1.M30 * f, m1.M31 * f, m1.M32 * f, m1.M33 * f,
	}
}

func (m1 Matrix) MulVector(v Vector) Vector {
	return Vector{
		m1.M00*v.X + m1.M01*v.Y + m1.M02*v.Z + m1.M03,
		m1.M10*v.X + m1.M11*v.Y + m1.M12*v.Z + m1.M13,
		m1.M20*v.X + m1.M21*v.Y + m1.M22*v.Z + m1.M23,
	}
}
