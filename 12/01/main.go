package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"

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

const NORTH = 0
const WEST = 1
const SOUTH = 2
const EAST = 3

type Ship struct {
	dir  int
	xLoc int
	yLoc int
}

type Command struct {
	name string
	val  int
}

func NewShip() Ship {
	return Ship{
		dir:  EAST,
		xLoc: 0,
		yLoc: 0,
	}
}

func CopyShip(s Ship) Ship {
	return Ship{
		dir:  s.dir,
		xLoc: s.xLoc,
		yLoc: s.yLoc,
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
		num, err := strconv.Atoi(res[2])
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
		ns.yLoc -= c.val
	case "S":
		ns.yLoc += c.val
	case "E":
		ns.xLoc += c.val
	case "W":
		ns.xLoc -= c.val
	case "F":
		switch ns.dir {
		case EAST:
			ns.xLoc += c.val
		case WEST:
			ns.xLoc -= c.val
		case SOUTH:
			ns.yLoc += c.val
		case NORTH:
			ns.yLoc -= c.val
		}
	case "R":
		amount := c.val / 90
		dir := ns.dir - amount
		if dir < 0 {
			dir += 4
		}
		ns.dir = dir
	case "L":
		amount := c.val / 90
		dir := ns.dir + amount
		if dir > 3 {
			dir -= 4
		}
		ns.dir = dir
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

	val := Abs(s.xLoc) + Abs(s.yLoc)
	fmt.Println(val)
}
