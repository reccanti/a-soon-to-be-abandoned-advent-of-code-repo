package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/reccanti/a-soon-to-be-abandoned-advent-of-code-repo/util"
)

// Our command types

type MaskCommand struct {
	ignoreMask int
	applyMask  int
}

type MemCommand struct {
	address int
	value   int
}

func parseMasks(str string) (*int, *int, error) {
	// create the "ignore" mask
	ignoreMaskStr := ""
	for _, char := range str {
		if string(char) == "X" {
			ignoreMaskStr = ignoreMaskStr + "1"
		} else {
			ignoreMaskStr = ignoreMaskStr + "0"
		}
	}
	ignoreMask, err := strconv.ParseInt(ignoreMaskStr, 2, 37)
	if err != nil {
		return nil, nil, errors.New("Unable to parse 'ignore' mask")
	}
	i := int(ignoreMask)

	// create the "apply" mask
	applyMaskStr := ""
	for _, char := range str {
		if string(char) == "X" {
			applyMaskStr = applyMaskStr + "0"
		} else {
			applyMaskStr = applyMaskStr + string(char)
		}
	}
	applyMask, err := strconv.ParseInt(applyMaskStr, 2, 37)
	if err != nil {
		return nil, nil, errors.New("Unable to parse 'apply' mask")
	}
	a := int(applyMask)

	// return our masks!
	return &i, &a, nil
}

// (mask|(mem)\[(\d+)\]) = (.+)
// Define a way to parse a string into a command
func parseCommand(str string) (interface{}, error) {
	exp := regexp.MustCompile(`(mask|(mem)\[(\d+)\]) = (.+)`)
	res := exp.FindStringSubmatch(str)
	if len(res) != 5 {
		return nil, errors.New("regex did not return the correct number of group matches")
	}
	// parse and return a "mask" command
	if res[1] == "mask" {
		i, a, err := parseMasks(res[4])
		if err != nil {
			return nil, err
		}
		cmd := MaskCommand{
			ignoreMask: *i,
			applyMask:  *a,
		}
		return cmd, nil
	}
	// parse and return a "mem" command
	if res[2] == "mem" {
		address, err := strconv.Atoi(res[3])
		if err != nil {
			return nil, err
		}
		value, err := strconv.Atoi(res[4])
		if err != nil {
			return nil, err
		}
		cmd := MemCommand{
			address: address,
			value:   value,
		}
		return cmd, nil
	}
	return nil, errors.New("Was not able to parse a command from the given input")
}

//  handle our state

type State struct {
	memory     map[int]int
	ignoreMask int
	applyMask  int
}

func Init() State {
	return State{
		memory:     map[int]int{},
		ignoreMask: 0,
		applyMask:  0,
	}
}

func (s State) Copy() State {
	memory := map[int]int{}
	for k, v := range s.memory {
		memory[k] = v
	}
	return State{
		memory:     memory,
		ignoreMask: s.ignoreMask,
		applyMask:  s.applyMask,
	}
}

func applyCommand(s State, cmd interface{}) State {
	ns := s.Copy()
	switch cmd.(type) {
	case MaskCommand:
		ns.ignoreMask = cmd.(MaskCommand).ignoreMask
		ns.applyMask = cmd.(MaskCommand).applyMask
	case MemCommand:
		// let's do some boolean math!
		val := cmd.(MemCommand).value&ns.ignoreMask | ns.applyMask
		ns.memory[cmd.(MemCommand).address] = val
	}
	return ns
}

func main() {
	// get the input
	filename := os.Args[1]
	lines, err := util.ParseRelativeFileSplit(filename, "\n")
	if err != nil {
		return
	}

	// create our list of commands
	cmds := []interface{}{}
	for _, line := range lines {
		cmd, err := parseCommand(line)
		if err != nil {
			fmt.Println(err)
			return
		}
		cmds = append(cmds, cmd)
	}

	// apply each command to the state
	s := Init()
	for _, cmd := range cmds {
		s = applyCommand(s, cmd)
	}

	// get the sum of all the values in memory
	sum := 0
	for _, val := range s.memory {
		sum += val
	}
	fmt.Println(sum)
}
