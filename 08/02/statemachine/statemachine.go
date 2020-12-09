package statemachine

import (
	"fmt"
)

type Instruction struct {
	Name  string
	Value int
}

type StateMachine struct {
	Index        int
	Acc          int
	Instructions []Instruction
	Finished     bool
}

func New(instructions []Instruction) StateMachine {
	s := StateMachine{
		Index:        0,
		Acc:          0,
		Instructions: instructions,
	}
	return s
}

func (s StateMachine) Next() StateMachine {
	if !s.Finished {
		ins := s.Instructions[s.Index]
		switch ins.Name {
		case "jmp":
			s.Index += ins.Value
		case "acc":
			s.Acc += ins.Value
			s.Index += 1
		default:
			s.Index += 1
		}
		if s.Index > len(s.Instructions) {
			s.Finished = true
		}
	}
	return s
}

func (s StateMachine) String() string {
	finished := "false"
	if s.Finished {
		finished = "true"
	}
	return fmt.Sprintf("Current State:\nindex: %d\nacc: %d\nfinished: %s", s.Index, s.Acc, finished)
}
