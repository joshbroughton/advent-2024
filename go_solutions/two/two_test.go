package main

import (
	"os"
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	byte_data, _ := os.ReadFile("test_input.txt")
	data := strings.Split(string(byte_data), "\n")
	result := Solve(data)
	if result != 2 {
		t.Fatalf(`Expected 2, got %q`, result)
	}
}

func TestSolveDampener(t *testing.T) {
	byte_data, _ := os.ReadFile("text_input_two.txt")
	data := strings.Split(string(byte_data), "\n")
	result := SolveDampener(data)
	if result != 10 {
		t.Fatalf(`Expected 10, got %v`, result)
	}
}

func TestIsSafe(t *testing.T) {
	data := []string{"1", "2", "4", "5"}
	result := IsSafe(data)
	if result != true {
		t.Fatalf(`Expected true, got %v`, result)
	}
}
