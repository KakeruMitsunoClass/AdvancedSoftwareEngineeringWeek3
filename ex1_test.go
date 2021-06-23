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

func TestIntToRome03(t *testing.T) {
	expected := "MCMXC"
	input := 1990
	result := intToRome(input)
	if expected != result {
		t.Errorf("Test03 is failed expected: %s, result: %s\n", expected, result)
	}
}

func TestIntToRome04(t *testing.T) {
	expected := "MMVIII"
	input := 2008
	result := intToRome(input)
	if expected != result {
		t.Errorf("Test04 is failed expected: %s, result: %s\n", expected, result)
	}
}

func TestIntToRome05(t *testing.T) {
	expected := "XCIX"
	input := 99
	result := intToRome(input)
	if expected != result {
		t.Errorf("Test05 is failed expected: %s, result: %s\n", expected, result)
	}
}

func TestIntToRome06(t *testing.T) {
	expected := "XLVII"
	input := 47
	result := intToRome(input)
	if expected != result {
		t.Errorf("Test06 is failed expected: %s, result: %s\n", expected, result)
	}
}
