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
	byte_data, _ := os.ReadFile("test_input.txt")
	data := strings.Split(string(byte_data), "\n")
	result := SolveDampener(data)
	if result != 4 {
		t.Fatalf(`Expected 4, got %q`, result)
	}
}
