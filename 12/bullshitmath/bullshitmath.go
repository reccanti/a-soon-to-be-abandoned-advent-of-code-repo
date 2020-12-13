/**
 * We're limiting ourselves to 2D here because I don't want
 * to deal with anything more complex
 */
package bullshitmath

import (
	"fmt"
	"math"
)

/**
 * | X Y |
 */
type Vector struct {
	vals []float64
}

func (v Vector) X() float64 {
	return v.vals[0]
}
func (v Vector) Y() float64 {
	return v.vals[1]
}

func MakeVector(x float64, y float64) Vector {
	return Vector{
		vals: []float64{x, y},
	}
}

// utility functions

func (v Vector) String() string {
	return fmt.Sprintf("| %v %v |", v.X(), v.Y())
}

func (v Vector) Copy() Vector {
	newVals := []float64{}
	for _, val := range v.vals {
		newVals = append(newVals, val)
	}
	return Vector{
		vals: newVals,
	}
}

/**
 * v1 + v2
 */
func VecAdd(v1 Vector, v2 Vector) Vector {
	x := v1.X() + v2.X()
	y := v1.Y() + v2.Y()
	return MakeVector(x, y)
}

/**
 * v1 * s
 */
func VecScale(v Vector, s float64) Vector {
	x := v.X() * s
	y := v.Y() * s
	return MakeVector(x, y)
}

/**
 * v1 - v2
 */
func VecSubtract(v1 Vector, v2 Vector) Vector {
	return VecAdd(v1, VecScale(v2, -1))
}

/**
 * Rotation functions
 */
func VecMultiplyMatrix(v Vector, m Matrix) Vector {
	x := v.X()*m.X1() + v.Y()*m.X2()
	y := v.X()*m.Y1() + v.Y()*m.Y2()
	// fmt.Println(x)
	// fmt.Println(y)
	return MakeVector(x, y)
}

func VecRotateCounterClockwise(v Vector, degrees float64) Vector {
	m := MakeRotationMatrix(degrees)
	// fmt.Println(m)
	return VecMultiplyMatrix(v, m)
}

func VecRotateClockwise(v Vector, degrees float64) Vector {
	return VecRotateCounterClockwise(v, -degrees)
}

/**
 * | X1 Y1 |
 * | X2 Y2 |
 */
type Matrix struct {
	vals []float64
}

// getters
func (m Matrix) X1() float64 {
	return m.vals[0]
}
func (m Matrix) Y1() float64 {
	return m.vals[1]
}
func (m Matrix) X2() float64 {
	return m.vals[2]
}
func (m Matrix) Y2() float64 {
	return m.vals[3]
}

func MakeMatrix(x1 float64, y1 float64, x2 float64, y2 float64) Matrix {
	return Matrix{
		vals: []float64{x1, y1, x2, y2},
	}
}

func MakeRotationMatrix(degrees float64) Matrix {
	rads := DegreesToRadians(degrees)
	// fmt.Println(rads)
	// fmt.Println(rads)
	return Matrix{
		vals: []float64{
			math.Cos(rads), -1 * math.Sin(rads),
			math.Sin(rads), math.Cos(rads),
		},
	}
}

// utility functions
func (m Matrix) String() string {
	return fmt.Sprintf("| %v %v |\n| %v %v |", m.X1(), m.Y1(), m.X2(), m.Y2())
}

func (m Matrix) Copy() Matrix {
	newVals := []float64{}
	for _, val := range m.vals {
		newVals = append(newVals, val)
	}
	return Matrix{
		vals: newVals,
	}
}

// random stuff

func DegreesToRadians(degree float64) float64 {
	return math.Pi * degree / 180
}
