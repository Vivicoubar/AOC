package main

import (
	_ "embed"
	"testing"
)

var inputTest = "input1.txt"
var ex = "ex.txt"

func TestExPart1(t *testing.T) {
	result := P1(ex)
	expected := 46
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1(t *testing.T) {
	result := P1(inputTest)
	expected := 7236
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		P1(inputTest)
	}
}

func TestExPart2(t *testing.T) {
	result := P2(ex)
	expected := 51
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := P2(inputTest)
	expected := 7521
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		P2(inputTest)
	}
}
