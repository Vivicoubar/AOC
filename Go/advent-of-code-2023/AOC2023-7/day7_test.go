package main

import (
	_ "embed"
	"testing"
)

var inputTest = "input1.txt"
var ex = "ex.txt"

func TestExPart1(t *testing.T) {
	result := P1(ex)
	expected := 6440
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1(t *testing.T) {
	result := P1(inputTest)
	expected := 248105065
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
	result := P2(ex)
	expected := 5905
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := P2(inputTest)
	expected := 249515436
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func BenchmarkPart2B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		P2(inputTest)
	}
}
