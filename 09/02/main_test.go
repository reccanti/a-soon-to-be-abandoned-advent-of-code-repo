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

	// this should fail because no 2 values will equal 10
	fail := checkSum(10, []int{1, 2, 3, 4, 5})
	if fail != false {
		t.Errorf("The given sum should fail")
	}
}
