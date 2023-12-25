package main

import (
	_ "embed"
	"testing"
)

var inputTest = "input1.txt"
var ex = "ex.txt"

func TestExPart1(t *testing.T) {
	result := P1(ex)
	expected := 54
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1(t *testing.T) {
	result := P1(inputTest)
	expected := 582590

	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		P1(inputTest)
	}
}
