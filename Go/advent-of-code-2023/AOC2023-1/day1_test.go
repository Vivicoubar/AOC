package main

import "testing"

var inputTest = "input1.txt"
var ex1 = "ex.txt"
var ex2 = "ex2.txt"

func TestExPart1(t *testing.T) {
	result := Part1(ex1)
	expected := 142
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestExPart2(t *testing.T) {
	result := Part2(ex2)
	expected := 281
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1(t *testing.T) {
	result := Part1(inputTest)
	expected := 56465
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(inputTest)
	expected := 55902
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
