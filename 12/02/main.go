package main

import (
	"errors"
	"fmt"
	// "math"
	"os"
	"regexp"
	"strconv"

	bs "github.com/reccanti/a-soon-to-be-abandoned-advent-of-code-repo/12/bullshitmath"
	"github.com/reccanti/a-soon-to-be-abandoned-advent-of-code-repo/util"
)

// Define our coordinates along these lines
// +----------------------+
// |          -3          |
// |          -2          |
// |          -1          |
// | -3 -2 -1  0 +1 +2 +3 |
// |          +1          |
// |          +2          |
// |          +3          |
// +----------------------+

var NORTH = bs.MakeVector(0, -1)
var WEST = bs.MakeVector(-1, 0)
var SOUTH = bs.MakeVector(0, 1)
var EAST = bs.MakeVector(1, 0)

type Ship struct {
	loc      bs.Vector
	waypoint bs.Vector
}

type Command struct {
	name string
	val  float64
}

func NewShip() Ship {
	return Ship{
		loc:      bs.MakeVector(0, 0),
		waypoint: bs.MakeVector(10, -1),
	}
}

func CopyShip(s Ship) Ship {
	return Ship{
		loc:      s.loc.Copy(),
		waypoint: s.waypoint.Copy(),
	}
}

func parseCommands(strs []string) ([]Command, error) {
	exp := regexp.MustCompile(`([FNSEWRL])(\d+)`)
	commands := []Command{}
	for _, val := range strs {
		res := exp.FindStringSubmatch(val)
		if len(res) != 3 {
			return nil, errors.New("command was not able to be parsed")
		}
		name := res[1]
		num, err := strconv.ParseFloat(res[2], 64)
		if err != nil {
			return nil, err
		}
		cmd := Command{
			name: name,
			val:  num,
		}
		commands = append(commands, cmd)
	}
	return commands, nil
}

func applyCommand(s Ship, c Command) Ship {
	ns := CopyShip(s)
	switch c.name {
	case "N":
		ns.waypoint = bs.VecAdd(ns.waypoint, bs.VecScale(NORTH, c.val))
	case "S":
		ns.waypoint = bs.VecAdd(ns.waypoint, bs.VecScale(SOUTH, c.val))
	case "E":
		ns.waypoint = bs.VecAdd(ns.waypoint, bs.VecScale(EAST, c.val))
	case "W":
		ns.waypoint = bs.VecAdd(ns.waypoint, bs.VecScale(WEST, c.val))
	case "R":
		locToWaypoint := bs.VecSubtract(ns.waypoint, ns.loc)
		rotatedWaypoint := bs.VecRotateClockwise(locToWaypoint, c.val)
		newWaypoint := bs.VecAdd(ns.loc, rotatedWaypoint)
		ns.waypoint = newWaypoint
	case "L":
		locToWaypoint := bs.VecSubtract(ns.waypoint, ns.loc)
		rotatedWaypoint := bs.VecRotateCounterClockwise(locToWaypoint, c.val)
		newWaypoint := bs.VecAdd(ns.loc, rotatedWaypoint)
		ns.waypoint = newWaypoint
	case "F":
		locToWaypoint := bs.VecSubtract(ns.waypoint, ns.loc)
		moveVector := bs.VecScale(locToWaypoint, c.val)
		ns.loc = bs.VecAdd(moveVector, ns.loc)
		ns.waypoint = bs.VecAdd(ns.loc, locToWaypoint)
	}
	return ns
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// get the input
	filename := os.Args[1]
	dirStrs, err := util.ParseRelativeFileSplit(filename, "\n")
	if err != nil {
		return
	}
	commands, err := parseCommands(dirStrs)
	if err != nil {
		return
	}

	// make a new ship
	s := NewShip()

	for _, val := range commands {
		s = applyCommand(s, val)
	}

	x := Abs(int(s.loc.X()))
	y := Abs(int(s.loc.Y()))
	fmt.Println(x + y)
}
