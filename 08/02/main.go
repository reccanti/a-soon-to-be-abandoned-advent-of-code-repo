package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/reccanti/a-soon-to-be-abandoned-advent-of-code-repo/08/02/statemachine"
	"github.com/reccanti/a-soon-to-be-abandoned-advent-of-code-repo/util"
)

// (.+) ([+-]\d+)
func parseInstructionString(str string) (*statemachine.Instruction, error) {
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
	return &statemachine.Instruction{Name: name, Value: value}, nil
}

/**
 * A convenient function for "flipping" our "jmp" and "nop" instructions
 */
func flip(i statemachine.Instruction) statemachine.Instruction {
	switch i.Name {
	case "jmp":
		return statemachine.Instruction{Name: "nop", Value: i.Value}
	case "nop":
		return statemachine.Instruction{Name: "jmp", Value: i.Value}
	default:
		return i
	}
}

type Node struct {
	state statemachine.StateMachine
	next  statemachine.Instruction
}

func getNext(s statemachine.StateMachine, in statemachine.Instruction) []Node {
	if in.Name == "jmp" || in.Name == "nop" {
		alt := flip(in)
		n1 := Node{state: s, next: in}
		n2 := Node{state: s, next: alt}
		return []Node{n1, n2}
	} else {
		n := Node{state: s, next: in}
		return []Node{n}
	}
}

func traverseAndFix(s statemachine.StateMachine, ins []statemachine.Instruction) (*statemachine.StateMachine, error) {
	// this will keep a record of all the states we've traversed
	// up until now. If we find a cycle, we'll revert to the previous state
	// and attempt to "flip" the instruction
	// toVisit := getNext(s, ins[s.Index])
	history := []statemachine.StateMachine{}

	// a lookup record for the different code lines we've visited up until
	// now. This is used to detect cycles
	visited := map[int]bool{}

	// use this to detect if we've already detected a cycle
	cycleDetected := false

	// cycle through the list until we pass the last instruction
	for s.Index < len(ins) {
		// if we've already visited this node, cycle back through our history
		// until we find a non-acc instruction. Once we've done that, we flip
		// the instruction and carry on
		//
		// This whole section is SUPER gross but it seems to work!
		if visited[s.Index] {
			if len(history) < 3 {
				return nil, errors.New("Unable to traverse the program without creating a cycle")
			}
			cycleDetected = true
			s = history[0]
			history = history[1:]
			for ins[s.Index].Name == "acc" {
				if len(history) < 3 {
					return nil, errors.New("Unable to traverse the program without creating a cycle")
				}
				s = history[0]
				history = history[1:]
			}
			in := flip(ins[s.Index])
			s = s.Execute(in)
		} else {
			visited[s.Index] = true
		}
		// add state to history list
		if !cycleDetected {
			// history = util.Prepend(s, history)
			history = append([]statemachine.StateMachine{s}, history...)
		}
		in := ins[s.Index]
		s = s.Execute(in)
	}

	return &s, nil
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
	instructions := []statemachine.Instruction{}
	for _, ins := range instructionStrings {
		instruction, err := parseInstructionString(ins)
		if err != nil {
			fmt.Println(err)
			return
		}
		instructions = append(instructions, *instruction)
	}

	s := statemachine.New()
	sp, err := traverseAndFix(s, instructions)
	if err != nil {
		fmt.Println(err)
	}
	s = *sp
	fmt.Println(s.Acc)
}
