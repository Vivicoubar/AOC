package main

import (
	_ "embed"
	"testing"
)

var inputTest = "input1.txt"
var ex = "ex.txt"

func TestExPart1(t *testing.T) {
	result := P1(ex)
	expected := 374
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1(t *testing.T) {
	result := P1(inputTest)
	expected := 10231178
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func BenchmarkPart1B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		P1(inputTest)
	}
}

func TestEx1Part2(t *testing.T) {
	result := P2(ex, 10)
	expected := 1030
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestEx2Part2(t *testing.T) {
	result := P2(ex, 100)
	expected := 8410
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := P2(inputTest, 1e6)
	expected := 622120986954
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func BenchmarkPart2B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		P2(inputTest, 1e6)
	}
}
