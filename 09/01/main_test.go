package main

import (
	"testing"
)

func TestCheckSum(t *testing.T) {
	// given a value and a list of preambles,
	// see if any of them add to the given sum

	// this should be true, because 5 + 4 = 9
	pass := checkSum(9, []int{1, 2, 3, 4, 5})
	if !pass {
		t.Errorf("The given sum is valid")
	}
}
