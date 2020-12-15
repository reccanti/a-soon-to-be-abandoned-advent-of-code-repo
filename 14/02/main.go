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
	branches   []int
}

type MemCommand struct {
	address int
	value   int
}

// parsing functions

func constructBranches(maskStr string) []int {
	// construct a "main" binary number from a given value.
	// if we encounter an "X" we log it as a pair of "floaters",
	// which will be a "0" and "1" variant of the given index
	base := 1
	main := 0
	floaters := [][]int{}
	for i := len(maskStr) - 1; i >= 0; i-- {
		char := string(maskStr[i])
		if char == "1" {
			main += 1 * base
		}
		if char == "X" {
			floaters = append(floaters, []int{0 * base, 1 * base})
		}
		base *= 2
	}

	// given our floaters from before, construct some
	// branching paths.
	branches := []int{0}
	for _, f := range floaters {
		newBranches := []int{}
		for _, b := range branches {
			newBranches = append(newBranches, b+f[0])
			newBranches = append(newBranches, b+f[1])
		}
		branches = newBranches
	}

	// construct all of the addresses
	addresses := []int{}
	for _, b := range branches {
		addresses = append(addresses, main+b)
	}

	return addresses
}

func constructIgnoreMask(maskStr string) int {
	ignoreMask := 0
	base := 1
	for i := len(maskStr) - 1; i >= 0; i-- {
		char := string(maskStr[i])
		if !(char == "X" || char == "1") {
			ignoreMask += 1 * base
		}
		base *= 2
	}
	return ignoreMask
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
		ignoreMask := constructIgnoreMask(res[4])
		branches := constructBranches(res[4])
		cmd := MaskCommand{
			ignoreMask: ignoreMask,
			branches:   branches,
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
	branches   []int
}

func Init() State {
	return State{
		memory:     map[int]int{},
		ignoreMask: 0,
		branches:   []int{},
	}
}

func (s State) Copy() State {
	memory := map[int]int{}
	for k, v := range s.memory {
		memory[k] = v
	}
	branches := []int{}
	for _, v := range s.branches {
		branches = append(branches, v)
	}
	return State{
		memory:     memory,
		ignoreMask: s.ignoreMask,
		branches:   branches,
	}
}

func applyCommand(s State, cmd interface{}) State {
	// fmt.Println(cmd)
	ns := s.Copy()
	switch cmd.(type) {
	case MaskCommand:
		ns.ignoreMask = cmd.(MaskCommand).ignoreMask
		ns.branches = cmd.(MaskCommand).branches
	case MemCommand:
		// let's do some boolean math!
		val := cmd.(MemCommand).value
		for _, b := range ns.branches {
			base := cmd.(MemCommand).address
			address := base&ns.ignoreMask | b
			ns.memory[address] = val
		}
	}
	return ns
}

// Heyo Here's the main loop!

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
