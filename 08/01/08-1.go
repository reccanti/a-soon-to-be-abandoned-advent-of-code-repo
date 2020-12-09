package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/reccanti/a-soon-to-be-abandoned-advent-of-code-repo/util"
)

type Instruction struct {
	name  string
	value int
}

type State struct {
	acc   int
	index int
}

// (.+) ([+-]\d+)
func parseInstructionString(str string) (*Instruction, error) {
	exp := regexp.MustCompile(`(jmp|nop|acc) ([+-]\d+)`)
	res := exp.FindStringSubmatch(str)
	if len(res) != 3 {
		return nil, errors.New("unable to parse the given instruction string")
	}
	name := res[1]
	value, err := strconv.Atoi(res[2])
	if err != nil {
		return nil, err
	}
	return &Instruction{name: name, value: value}, nil
}

/**
 * Traverse through each of our instructions, halting if we
 * ever hit the same instruction a second time
 */
func traverseAndHalt(instructions []Instruction) State {
	state := State{
		index: 0,
		acc:   0,
	}
	visited := map[int]bool{}
	for state.index < len(instructions) && state.index >= 0 {
		// if we've already visited this instruction, break out of our traversal
		// otherwise add it to our "visited" list
		if visited[state.index] {
			break
		}
		visited[state.index] = true

		// update the state based on the action
		ins := instructions[state.index]
		fmt.Println(ins)
		switch ins.name {
		case "jmp":
			state.index += ins.value
		case "acc":
			state.acc += ins.value
			state.index += 1
		default:
			state.index += 1
		}
	}
	return state
}

func main() {
	// parse our input
	filename := os.Args[1]
	input, err := util.ParseRelativeFile(filename)
	if err != nil {
		return
	}
	instructionBlock := *input
	instructionStrings := strings.Split(instructionBlock, "\n")
	instructions := []Instruction{}
	for _, ins := range instructionStrings {
		instruction, err := parseInstructionString(ins)
		if err != nil {
			fmt.Println(err)
			return
		}
		instructions = append(instructions, *instruction)
	}

	// let's traverse those instructions!
	state := traverseAndHalt(instructions)
	fmt.Println(state.acc)
}
