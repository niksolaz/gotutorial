package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{-5, -3},
		{99999, 100001},
	}
	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}

func TestCalculateMultiply(t *testing.T) {
	var mults = []struct {
		input    int
		expected int
	}{
		{2, 6},
		{-1, -3},
		{0, 0},
		{-5, -15},
		{90, 270},
	}

	for _, mult := range mults {
		if outp := CalculateMultiply(mult.input); outp != mult.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", mult.input, mult.expected, outp)
		}
	}
}

func TestPrintString(t *testing.T) {
	PrintString("Nicola")
}
