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

	s := statemachine.New(instructions)

	statemachine.Next(s)
	fmt.Println(s.ToString())

}
