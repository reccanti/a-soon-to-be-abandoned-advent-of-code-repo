package main

import (
	"testing"
)

func TestParseBagString(t *testing.T) {
	bagString := "light red bags contain 1 bright white bag, 2 muted yellow bags."
	container, containing, err := parseBagString(bagString)
	if err != nil {
		t.Errorf("should properly parse bag string")
	}
	if container != "light red" {
		t.Errorf("'container' improperly parsed")
	}
	if len(containing) != 2 || containing["bright white"] != 1 || containing["muted yellow"] != 2 {
		t.Errorf("'containing' map improperly parsed")
	}
}
