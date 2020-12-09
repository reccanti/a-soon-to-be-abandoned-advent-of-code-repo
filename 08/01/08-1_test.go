package main

import (
	"testing"
)

func TestParseInstructionString(t *testing.T) {
	_, err := parseInstructionString("unparseable")
	if err == nil {
		t.Errorf("the given string should not be parseable")
	}

	_, err = parseInstructionString("jmp val")
	if err == nil {
		t.Errorf("should not be able to parse instruction without numeric argument")
	}

	_, err = parseInstructionString("som +5")
	if err == nil {
		t.Errorf("should not be able to parse invalid instruction")
	}

	ins, invalidErr := parseInstructionString("jmp -4")
	if invalidErr != nil {
		t.Errorf("should be able to parse valid instruction")
	}
	if ins.name != "jmp" {
		t.Errorf("name parsed incorrectly")
	}
	if ins.value != -4 {
		t.Errorf("value parsed incorrectly")
	}
}
