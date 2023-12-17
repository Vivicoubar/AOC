package main

import (
	_ "embed"
	"testing"
)

var inputTest = "input1.txt"
var ex1 = "ex.txt"
var ex2 = "ex2.txt"
var ex3 = "ex3.txt"

func TestExPart1(t *testing.T) {
	result := P1(ex1)
	expected := 2
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
func TestEx2Part1(t *testing.T) {
	result := P1(ex2)
	expected := 6
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1(t *testing.T) {
	result := P1(inputTest)
	expected := 19099
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func BenchmarkPart1B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		P1(inputTest)
	}
}

func TestExPart2(t *testing.T) {
	result := P2(ex3)
	expected := 6
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := P2(inputTest)
	expected := 17099847107071
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func BenchmarkPart2B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		P2(inputTest)
	}
}
