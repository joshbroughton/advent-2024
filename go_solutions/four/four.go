package main

import (
	"fmt"
	"os"
	"strings"
	"slices"
)

func main() {
	byte_data, _ := os.ReadFile("../inputs/four.txt")
	// byte_data, _ := os.ReadFile("test_input.txt")
	data := strings.Split(string(byte_data), "\n")
	fmt.Println(Solve(data))
}

func Solve(rows []string) int {
	columns := BuildColumns(rows)
	diagonals := BuildDiagonals(rows)
	reversed_rows := make([]string, len(rows))
	copy(reversed_rows, rows)
	slices.Reverse(reversed_rows)
	reversed_rows = reversed_rows[1:]
	more_diagonals := BuildDiagonals(reversed_rows)
	words := [][]string{rows, columns, diagonals, more_diagonals}

	return (CountXmas(words))
}

func BuildColumns(rows []string) []string {
	column_number := len(rows)-1
	columns := []string{}
	for i:=0;i<column_number;i++ {
		var sb strings.Builder
		for j:=0;j<len(rows[i]);j++ {
			sb.WriteString(string(rows[j][i]))
		}
		columns = append(columns, sb.String())
	}

	return columns
}

func BuildDiagonals(rows []string) []string {
	diagonals := []string{}
	for i:=0;i<len(rows)-1;i++ {
		var sb strings.Builder
		var sbb strings.Builder
		k := 0
		x := len(rows[i])-1
		temp := i
		for j:=i;j>=0;j-- {
			sb.WriteString(string(rows[temp][k]))
			sbb.WriteString(string(rows[temp][x]))
			k++
			x--
			temp--
		}
		diagonals = append(diagonals, sb.String())
		diagonals = append(diagonals, sbb.String())
	}

	return diagonals
}

func CountXmas(words [][]string) int {
	count := 0
	for i:=0;i<len(words);i++ {
		for j:=0;j<len(words[i]);j++ {
			count = count + strings.Count(words[i][j], "XMAS")
			count = count + strings.Count(words[i][j], "SAMX")
		}
	}

	return count
}

