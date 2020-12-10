package statemachine

import (
	"fmt"
)

type Instruction struct {
	Name  string
	Value int
}

type StateMachine struct {
	Index int
	Acc   int
	// Finished bool
}

func New() StateMachine {
	s := StateMachine{
		Index: 0,
		Acc:   0,
	}
	return s
}

// func (s StateMachine) Next(instructions []Instruction) StateMachine {
// 	if !s.Finished {
// 		ins := instructions[s.Index]
// 		s = s.Execute(ins)
// 		if s.Index > len(instructions) {
// 			s.Finished = true
// 		}
// 	}
// 	return s
// }

func (s StateMachine) Execute(ins Instruction) StateMachine {
	switch ins.Name {
	case "jmp":
		s.Index += ins.Value
	case "acc":
		s.Acc += ins.Value
		s.Index += 1
	default:
		s.Index += 1
	}
	return s
}

func (s StateMachine) String() string {
	// finished := "false"
	// if s.Finished {
	// 	finished = "true"
	// }
	return fmt.Sprintf("Current State:\nindex: %d\nacc: %d", s.Index, s.Acc)
}
