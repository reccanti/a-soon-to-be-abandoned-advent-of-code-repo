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
 * Returns 2 values, the index at which the cycle was detected, and whether
 * or not the state machine finished. If the state machine finished, there is
 * no cycle, so we return an index of -1
 */
func detectCycle(s1 statemachine.StateMachine, s2 statemachine.StateMachine) (int, bool) {
	// iterate through each of these state machines until they
	// either finish or reach a common index
	for {
		s1 = s1.Next()
		s2 = s2.Next().Next()

		if s1.Finished {
			fmt.Println("First state machine finished with acc", s1.Acc)
			return -1, true
		}
		if s2.Finished {
			fmt.Println("Second state machine finished with acc", s2.Acc)
			return -1, true
		}
		if s1.Index == s2.Index {
			fmt.Println("Cycle detected at index", s1.Index)
			return s1.Index, false
		}
	}
}

type InstructionLog struct {
	lineNumber  int
	instruction statemachine.Instruction
}

func getCycleInstructions(s statemachine.StateMachine, cycleIndex int) []InstructionLog {
	// get to the index at which the cycle occurs
	for s.Index != cycleIndex {
		s = s.Next()
	}

	// once we get to that state, start logging the instructions
	initLog := InstructionLog{
		lineNumber:  s.Index,
		instruction: s.Instructions[s.Index],
	}
	instructions := []InstructionLog{initLog}
	s = s.Next()
	for s.Index != cycleIndex {
		log := InstructionLog{
			lineNumber:  s.Index,
			instruction: s.Instructions[s.Index],
		}
		instructions = append(instructions, log)
		s = s.Next()
	}
	return instructions
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

	s1 := statemachine.New(instructions)
	s2 := statemachine.New(instructions)

	// see if we can detect a cycle
	index, didFinish := detectCycle(s1, s2)
	if didFinish {
		fmt.Println(fmt.Errorf("Did not detect a cycle"))
		return
	}

	// get the instructions in that cycle
	logs := getCycleInstructions(s1, index)
	filteredLogs := []InstructionLog{}
	for _, log := range logs {
		if log.instruction.Name != "acc" {
			filteredLogs = append(filteredLogs, log)
		}
	}

	fmt.Println(filteredLogs)

}
