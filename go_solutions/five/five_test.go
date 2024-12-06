package main

import (
	"fmt"
	"testing"
	"os"
	"strings"
)

// func TestParseInput(t *testing.T) {
// 	byte_data, _ := os.ReadFile("test_input.txt")
// 	data := strings.Split(string(byte_data), "\n")
// 	rules, commands := ParseInput(data)
// 	fmt.Println(rules)
// 	fmt.Println(commands)
// }

func TestSolution(t *testing.T) {
	byte_data, _ := os.ReadFile("test_input.txt")
	data := strings.Split(string(byte_data), "\n")
	rules, commands := ParseInput(data)
	valid_commands := Solve(rules, commands)
	fmt.Println(valid_commands)
	result := SumMiddles(valid_commands)
	if result != 143 {
		t.Fatalf(`Expected 143, got %q`, result)
	}
}
