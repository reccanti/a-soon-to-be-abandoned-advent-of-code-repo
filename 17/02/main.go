package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/reccanti/a-soon-to-be-abandoned-advent-of-code-repo/util"
)

// CUBE SHIT!!!

type Cube struct {
	coordinates []int
}

func NewCube(coordinates ...int) Cube {
	return Cube{
		coordinates,
	}
}

// number of neighbors: 3^n-1

// 2D cube; 3^2-1; 8; (0, 0): (-1, -1) (-1, 0) (-1, 1) (0, -1) (0, 1) (1, -1) (1, 0) (1, 1)
// 3D cube; 3^3-1; 26
// 4D cube; 3^4-1; 80
func (c Cube) GetNeighbors() []Cube {
	// This works by building a compounding series of paths.
	// Each iteration of the loop, we take all of our existing
	// paths and append one of the immediately surrounding neighbors.
	// This exponentially increases the number of paths
	paths := [][]int{[]int{}}
	for _, cur := range c.coordinates {
		newPaths := [][]int{}
		for i := -1; i <= 1; i++ {
			for _, p := range paths {
				// copy the current path
				cpy := make([]int, len(p))
				copy(cpy, p)

				// append our value to this path and add that to the new paths
				cpy = append(cpy, cur+i)
				newPaths = append(newPaths, cpy)
			}
		}
		paths = newPaths
	}

	// Once we've constructed the paths, use them to construct
	// new cubes
	neighbors := []Cube{}
	for _, p := range paths {
		isSame := true
		for i, coord := range c.coordinates {
			if coord != p[i] {
				isSame = false
				break
			}
		}
		if !isSame {
			c := NewCube(p...)
			neighbors = append(neighbors, c)
		}
	}
	return neighbors
}

type Coordinates struct {
	next map[int]Coordinates
}

func (c Coordinates) Copy() Coordinates {
	newCoords := Coordinates{
		next: map[int]Coordinates{},
	}
	for k, v := range c.next {
		newCoords.next[k] = v.Copy()
	}
	return newCoords
}

type CubeSpace struct {
	cubes       []Cube
	coordinates Coordinates
}

func NewSpace() CubeSpace {
	return CubeSpace{
		cubes: []Cube{},
		coordinates: Coordinates{
			next: map[int]Coordinates{},
		},
	}
}

func (s CubeSpace) HasCube(c Cube) bool {
	// _, hasX := s.coordinates[c.x]
	// if hasX {
	// 	_, hasY := s.coordinates[c.x][c.y]
	// 	if hasY {
	// 		_, hasZ := s.coordinates[c.x][c.y][c.z]
	// 		if hasZ {
	// 			return true
	// 		}
	// 	}
	// }
	// return false
	cur := s.coordinates.next
	hasCube := true
	for _, coord := range c.coordinates {
		_, hasEntry := cur[coord]
		if !hasEntry {
			hasCube = false
			break
		}
		cur = cur[coord].next
	}
	return hasCube
}

func (s CubeSpace) AddCube(c Cube) CubeSpace {
	if !s.HasCube(c) {
		ns := NewSpace()
		// update list of cubes
		ns.cubes = append(s.cubes, c)

		// update coordinates
		ns.coordinates = s.coordinates.Copy()
		cur := ns.coordinates.next
		for _, coord := range c.coordinates {
			_, hasEntry := cur[coord]
			if !hasEntry {
				cur[coord] = Coordinates{
					next: map[int]Coordinates{},
				}
			}
			cur = cur[coord].next
		}
		return ns
	} else {
		return s
	}
}

func (s CubeSpace) AddCubes(cs []Cube) CubeSpace {
	for _, c := range cs {
		s = s.AddCube(c)
	}
	return s
}

func (s CubeSpace) GetActiveNeighbors(c Cube) []Cube {
	ns := c.GetNeighbors()
	active := []Cube{}
	for _, n := range ns {
		if s.HasCube(n) {
			active = append(active, n)
		}
	}
	return active
}

func (s CubeSpace) Next() CubeSpace {
	nextCubes := []Cube{}

	for _, c := range s.cubes {
		ns := s.GetActiveNeighbors(c)

		// Step 1: Determine whether to carry-over occupied cubes
		if len(ns) >= 2 && len(ns) <= 3 {
			nextCubes = append(nextCubes, c)
		}

		// Step 2: Determine whether to flip unoccupied cubes
		allneighbors := c.GetNeighbors()
		for _, n := range allneighbors {
			if !s.HasCube(n) {
				nns := s.GetActiveNeighbors(n)
				// if n.coordinates[0] == 0 && n.coordinates[1] == 0 && n.coordinates[2] == -1 && n.coordinates[3] == -1 {
				// 	fmt.Println(s.GetActiveNeighbors(n))
				// }
				if len(nns) == 3 {
					nextCubes = append(nextCubes, n)
				}
			}

		}
	}
	// fmt.Println(len(nextCubes))

	newspace := NewSpace()
	newspace = newspace.AddCubes(nextCubes)
	return newspace
}

// Stuff to parse our input!

func parseCubes(str string) CubeSpace {
	lines := strings.Split(str, "\n")
	cubes := []Cube{}
	for r, l := range lines {
		for c, char := range l {
			if string(char) == "#" {
				c := Cube{
					coordinates: []int{c, r, 0, 0},
				}
				cubes = append(cubes, c)
			}
		}
	}

	s := NewSpace()
	s = s.AddCubes(cubes)
	return s
}

func main() {
	// c := NewCube(0, 0, 0, 0)
	// fmt.Println(c)
	// neighbors := c.GetNeighbors()
	// fmt.Println(len(neighbors))

	// space := NewSpace()
	// space2 := space.AddCube(NewCube(0, 0, 0, 0))

	// fmt.Println(space)
	// fmt.Println(space2)

	// get the input
	filename := os.Args[1]
	input, err := util.ParseRelativeFile(filename)
	if err != nil {
		return
	}

	space := parseCubes(*input)
	for i := 1; i <= 6; i++ {
		space = space.Next()
	}
	fmt.Println(len(space.cubes))
}
