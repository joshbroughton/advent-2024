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
		faults := 0
		first := true
		length := len(levels) - 1
		for j := 0; j < length; j++ {
			current, _ := strconv.Atoi(levels[j])
			next, _ := strconv.Atoi(levels[j+1])
			if direction {
				if current >= next || next-current > 3 {
					faults++
					if first {
						next, _ = strconv.Atoi(levels[j+2])
						if current >= next || next-current > 3 {
							levels = levels[1:]
							direction = Direction(levels[0], levels[1])
						} else {
							levels = append(levels[:j+1], levels[j+2:]...)
						}
					} else {
						levels = append(levels[:j+1], levels[j+2:]...)
					}
					length--
					j--
				}
			} else {
				if current <= next || current-next > 3 {
					faults++
					if first {
						next, _ = strconv.Atoi(levels[j+2])
						if current <= next || current-next > 3 {
							levels = levels[1:]
							direction = Direction(levels[0], levels[1])
						} else {
							levels = append(levels[:j+1], levels[j+2:]...)
						}
					} else {
						levels = append(levels[:j+1], levels[j+2:]...)
					}
					length--
					j--
				}
			}
			if faults > 1 {
				break
			}
			first = false
		}
		if faults <= 1 {
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
