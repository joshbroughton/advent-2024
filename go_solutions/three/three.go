package main

import (
	"os"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	byte_data, _ := os.ReadFile("../inputs/three.txt")
	// byte_data, _ := os.ReadFile("test_input.txt")
	data := string(byte_data)
	fmt.Println(SolveConditional(data))
}

func Solve(data string) int {
	commands := Match(data)
	sum := 0
	for i:=0;i<len(commands);i++ {
		sum = sum + RunCommand(commands[i])
	}

	return sum
}

func SolveConditional(data string) int {
	commands := MatchConditional(data)
	sum := 0
	run := true
	for i:=0;i<len(commands);i++ {
		command := commands[i]
		fmt.Println(command)
		if command == "do()" {
			run = true
		} else if command == "don't()" {
			run = false
		} else if run == true {
			sum = sum + RunCommand(commands[i])
		}
	}

	return sum
}

func Match(input string) []string {
	r, _ := regexp.Compile(`mul\(\d+\,\d+\)`)
	return r.FindAllString(input, -1)
}

func MatchConditional(input string) []string {
	r, _ := regexp.Compile(`mul\(\d+\,\d+\)|don?\'?t?\(\)`)
	return r.FindAllString(input, -1)
}

func RunCommand(command string) int {
	r, _ := regexp.Compile(`\d+`)
	digits := r.FindAllString(command, -1)
	x, _ := strconv.Atoi(digits[0])
	y, _ := strconv.Atoi(digits[1])

	return x * y
}
