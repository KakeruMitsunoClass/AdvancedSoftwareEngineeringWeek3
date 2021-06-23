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

func TestIntToRome02(t *testing.T) {
	expected := "CM"
	input := 900
	result := intToRome(input)
	if expected != result {
		t.Errorf("Test02 is failed expected: %s, result: %s\n", expected, result)
	}
}
