package statemachine

import (
	"fmt"
)

type Instruction struct {
	Name  string
	Value int
}

type StateMachine struct {
	index        int
	acc          int
	instructions []Instruction
}

func New(instructions []Instruction) StateMachine {
	s := StateMachine{
		index:        0,
		acc:          0,
		instructions: instructions,
	}
	return s
}

func Next(s StateMachine) {
	ins := s.instructions[s.index]
	switch ins.Name {
	case "jmp":
		s.index += ins.Value
	case "acc":
		s.acc += ins.Value
		s.index += 1
	default:
		s.index += 1
	}
}

func ToString(s StateMachine) string {
	return fmt.Sprintf("Current State:\nindex: %d\nacc: %d", s.index, s.acc)
}
