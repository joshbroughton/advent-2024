package main

import (
	"os"
	"strings"
	"strconv"
	"fmt"
)

func main() {
	// byte_data, _ := os.ReadFile("../inputs/four.txt")
	byte_data, _ := os.ReadFile("test_input.txt")
	data := strings.Split(string(byte_data), "\n")
	fmt.Println(data)
}

// create map for banned order and list of commands
func ParseInput(data []string) (map[string]string, []string) {
	change := true
	rules := make(map[string]string)
	commands := []string{}
	for i:=0;i<len(data);i++ {
		if data[i] == "" {
			change = false
			continue
		}
		if change {
			key := data[i][3:5]
			value := data[i][0:2]
			val, ok := rules[key]
			if ok {
				rules[key] = val + value
			} else {
				rules[key] = value
			}
		} else {
			commands = append(commands, data[i])
		}
	}

	return rules, commands
}

func Solve(rules map[string]string, commands []string) [][]string {
	valid_commands := [][]string{}
	for i:=0;i<len(commands);i++ {
		valid := true
		row := strings.Split(commands[i], ",")
		numbers := []string{}
		for j:=0;j<len(row);j++ {
			illegals := rules[row[j]]
			for k:=0;k<len(numbers);k++ {
				// leaving this as a string was a mistake it creates new illegal values
				if strings.Contains(illegals, numbers[k]) {
					valid = false
					break
				}
			}
			numbers = append(numbers, row[j])
			if !valid {
				break
			}
		}
		if valid {
			valid_commands = append(valid_commands, row)
		}
	}

	return valid_commands
}

func SumMiddles(commands [][]string) int {
	middle_sum := 0
	fmt.Println(commands)
	for i:=0;i<len(commands);i++ {
		middle_index := len(commands[i]) / 2
		fmt.Println(commands[i][middle_index])
		middle_value, _ := strconv.Atoi(commands[i][middle_index])
		middle_sum = middle_sum + middle_value
	}

	return middle_sum
}
