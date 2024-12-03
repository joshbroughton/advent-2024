package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	byte_data, _ := os.ReadFile("../inputs/two.txt")
	// byte_data, _ := os.ReadFile("test_input.txt")
	data := strings.Split(string(byte_data), "\n")
	fmt.Print(Solve(data))
}

func Solve(reports []string) int {
	safe := 0
	for i := 0; i < len(reports)-1; i++ {
		levels := strings.Fields(reports[i])
		direction := Direction(levels[0], levels[1])
		is_safe := true
		for j := 0; j < len(levels)-1; j++ {
			current, _ := strconv.Atoi(levels[j])
			next, _ := strconv.Atoi(levels[j+1])
			if direction {
				if current >= next || next-current > 3 {
					is_safe = false
					break
				}
			} else {
				if current <= next || current-next > 3 {
					is_safe = false
					break
				}
			}
		}
		if is_safe {
			safe++
		}
	}

	return safe
}

func SolveDampener(reports []string) int {
	safe := 0
	for i := 0; i < len(reports)-1; i++ {
		levels := strings.Fields(reports[i])
		direction := Direction(levels[0], levels[1])
		is_safe := true
		for j := 0; j < len(levels)-1; j++ {
			current, _ := strconv.Atoi(levels[j])
			next, _ := strconv.Atoi(levels[j+1])
			if direction {
				if current >= next || next-current > 3 {
					is_safe = false
					break
				}
			} else {
				if current <= next || current-next > 3 {
					is_safe = false
					break
				}
			}
		}
		if is_safe {
			safe++
		}
	}

	return safe
}

// true for asc, false for desc
func Direction(first_string string, second_string string) bool {
	first, _ := strconv.Atoi(first_string)
	second, _ := strconv.Atoi(second_string)
	if first > second {
		return false
	} else {
		return true
	}
}
