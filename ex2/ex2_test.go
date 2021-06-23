package myPackage

import "testing"

func TestRomeToInt01(t *testing.T) {
	expected := 4000
	input := "MMMM"
	result := romaToInt(input)
	if expected != result {
		t.Errorf("Test01 is failed expected: %d, result: %d\n", expected, result)
	}
}

func TestRomeToInt02(t *testing.T) {
	expected := 2008
	input := "MMVIII"
	result := romaToInt(input)
	if expected != result {
		t.Errorf("Test02 is failed expected: %d, result: %d\n", expected, result)
	}
}

func TestRomeToInt03(t *testing.T) {
	expected := 99
	input := "XCIX"
	result := romaToInt(input)
	if expected != result {
		t.Errorf("Test03 is failed expected: %d, result: %d\n", expected, result)
	}
}

func TestRomeToInt04(t *testing.T) {
	expected := 47
	input := "XLVII"
	result := romaToInt(input)
	if expected != result {
		t.Errorf("Test04 is failed expected: %d, result: %d\n", expected, result)
	}
}

func TestRomeToInt05(t *testing.T) {
	expected := 1990
	input := "MCMXC"
	result := romaToInt(input)
	if expected != result {
		t.Errorf("Test05 is failed expected: %d, result: %d\n", expected, result)
	}
}
