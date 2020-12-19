package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/reccanti/a-soon-to-be-abandoned-advent-of-code-repo/util"
)

// CUBE SHIT!!!

type Cube struct {
	x int
	y int
	z int
}

func (c Cube) GetNeighbors() []Cube {
	neighbors := []Cube{}
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				if !(i == 0 && j == 0 && k == 0) {
					neighbor := Cube{
						x: c.x + i,
						y: c.y + j,
						z: c.z + k,
					}
					neighbors = append(neighbors, neighbor)
				}
			}
		}
	}
	return neighbors
}

type CubeSpace struct {
	cubes       []Cube
	coordinates map[int]map[int]map[int]bool // this is a garbage data structure. I'm sorry
}

func NewSpace() CubeSpace {
	return CubeSpace{
		cubes:       []Cube{},
		coordinates: map[int]map[int]map[int]bool{},
	}
}

func (s CubeSpace) HasCube(c Cube) bool {
	_, hasX := s.coordinates[c.x]
	if hasX {
		_, hasY := s.coordinates[c.x][c.y]
		if hasY {
			_, hasZ := s.coordinates[c.x][c.y][c.z]
			if hasZ {
				return true
			}
		}
	}
	return false
}

func (s CubeSpace) AddCube(c Cube) CubeSpace {
	if !s.HasCube(c) {
		// update list of cubes
		s.cubes = append(s.cubes, c)

		// update coordinates
		_, hasX := s.coordinates[c.x]
		if !hasX {
			s.coordinates[c.x] = map[int]map[int]bool{}
		}
		_, hasY := s.coordinates[c.x][c.y]
		if !hasY {
			s.coordinates[c.x][c.y] = map[int]bool{}
		}
		_, hasZ := s.coordinates[c.x][c.y][c.z]
		if !hasZ {
			s.coordinates[c.x][c.y][c.z] = false
		}
		s.coordinates[c.x][c.y][c.z] = true
	}
	return s
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
				if len(nns) == 3 {
					nextCubes = append(nextCubes, n)
				}
			}

		}
	}

	// fmt.Println(nextCubes)

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
					x: c,
					y: r,
					z: 0,
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
	// get the input
	filename := os.Args[1]
	input, err := util.ParseRelativeFile(filename)
	if err != nil {
		return
	}

	space := parseCubes(*input)
	space = space.Next()
	fmt.Println(space.cubes)
	// fmt.Println(space.cubes)
	// for i := 1; i <= 6; i++ {
	// 	space = space.Next()
	// 	fmt.Println(space.cubes)
	// }
	fmt.Println(len(space.cubes))
}
