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
