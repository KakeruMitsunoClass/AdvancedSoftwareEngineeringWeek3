package main

import "testing"

func TestIntToRome01(t *testing.T) {
	expected := "MMMM"
	input := 4000
	result := intToRome(input)
	if expected != result {
		t.Errorf("Test01 is failed expected: %s, result: %s\n", expected, result)
	}
}
