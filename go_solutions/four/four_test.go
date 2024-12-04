package main

import (
	"os"
	"strings"
	"testing"
	"fmt"
	"slices"
)

func TestSolve(t *testing.T) {
	byte_data, _ := os.ReadFile("test_input.txt")
	data := strings.Split(string(byte_data), "\n")
	result := Solve(data)
	if result != 18 {
		t.Fatalf(`Expected 18, got %q`, result)
	}
}

// func TestBuildColumns(t *testing.T) {
// 	byte_data, _ := os.ReadFile("test_input.txt")
// 	data := strings.Split(string(byte_data), "\n")
// 	result := BuildColumns(data)
// 	fmt.Println(result)
// }

func TestBuildDiagonals(t *testing.T) {
	byte_data, _ := os.ReadFile("test_input.txt")
	data := strings.Split(string(byte_data), "\n")
	slices.Reverse(data)
	data = data[1:]
	result := BuildDiagonals(data)
	fmt.Println(result)
}
