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
	fmt.Println(Solve(data))
	fmt.Println(SolveDampener(data))
}

func Solve(reports []string) int {
	safe := 0
	for i := 0; i < len(reports)-1; i++ {
		report := strings.Fields(reports[i])
		if IsSafe(report) {
			safe++
		}
	}

	return safe
}

func SolveDampener(reports []string) int {
	safe := 0
	for i := 0; i < len(reports)-1; i++ {
		report := strings.Fields(reports[i])
		if IsSafe(report) {
			safe++
		} else {
			for k := 0; k < len(report); k++ {
				test_report := make([]string, len(report))
				copy(test_report, report)
				test_report = append(test_report[:k], test_report[k+1:]...)
				if IsSafe(test_report) {
					safe++
					break
				}
			}
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

func IsSafe(report []string) bool {
	report_safe := true
	direction := Direction(report[0], report[1])
	for j := 0; j < len(report)-1; j++ {
		current, _ := strconv.Atoi(report[j])
		next, _ := strconv.Atoi(report[j+1])
		if direction {
			if current >= next || next-current > 3 {
				report_safe = false
				break
			}
		} else {
			if current <= next || current-next > 3 {
				report_safe = false
				break
			}
		}
	}

	return report_safe
}
