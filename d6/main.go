package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type operators int

const (
	ADD operators = iota
	MULTIPLY
)

func (o operators) String() string {
	switch o {
	case ADD:
		return "+"
	case MULTIPLY:
		return "*"
	default:
		return "unknown"
	}
}

func main() {
	rows, ops := readInput("./input_test")

	// Transpose: convert rows to columns
	numCols := len(rows[0])
	cols := make([][]int, numCols)

	for colIdx := 0; colIdx < numCols; colIdx++ {
		cols[colIdx] = make([]int, len(rows))
		for rowIdx := 0; rowIdx < len(rows); rowIdx++ {
			cols[colIdx][rowIdx] = rows[rowIdx][colIdx]
		}
	}

	// Calculate sum
	sum := 0
	for i, col := range cols {
		fmt.Printf("col %d: %v, op: %s\n", i, col, ops[i])

		// Start with the first value
		result := col[0]

		// Apply operator to remaining values
		mathOps := ops[i]
		for j := 1; j < len(col); j++ {
			switch mathOps {
			case ADD:
				result += col[j]
			case MULTIPLY:
				result *= col[j]
			}
		}

		fmt.Printf("result: %d\n", result)
		sum += result
	}

	fmt.Println("part1: ", sum)
}

func readInput(path string) ([][]int, []operators) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)

	rows := make([][]int, 0)
	var ops []operators

	for s.Scan() {
		rowStr := strings.TrimSpace(s.Text())
		fields := strings.Fields(rowStr)

		currentRow := make([]int, 0)

		for _, colElement := range fields {
			num, err := strconv.Atoi(colElement)
			if err != nil {
				switch colElement {
				case "+":
					ops = append(ops, ADD)
				case "*":
					ops = append(ops, MULTIPLY)
				}
				continue
			}
			currentRow = append(currentRow, num)
		}

		if len(currentRow) > 0 {
			rows = append(rows, currentRow)
		}
	}
	return rows, ops
}
