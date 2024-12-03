package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	byte_data, _ := os.ReadFile("../inputs/one.txt")
	data := strings.Fields(string(byte_data))
	left := []int{}
	right := []int{}
	// create an array for each column
	for i := 0; i < len(data); i++ {
		num, _ := strconv.Atoi(data[i])
		if i%2 == 0 {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}

	left = sort(left)
	right = sort(right)

	// sum := 0
	// for i:=0; i<len(left); i++ {
	// 	sum = sum + int(math.Abs(float64(left[i] - right[i])))
	// }
	// fmt.Print(sum)
	sim_score := 0
	for i := 0; i < len(left); i++ {
		for j := 0; j < len(right); j++ {
			if left[i] == right[j] {
				sim_score = sim_score + right[j]
			}
		}
	}
	fmt.Print(sim_score)
}

func sort(input []int) []int {
	output := []int{}
	for {
		least := input[0]
		least_index := 0
		for j := 1; j < len(input); j++ {
			if input[j] < least {
				least = input[j]
				least_index = j
			}
		}
		output = append(output, least)
		input = append(input[:least_index], input[least_index+1:]...)
		if len(input) == 0 {
			break
		}
	}

	return output
}
