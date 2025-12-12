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
	// day1()
	rows, ops := readInputDay2("./input_test")
	fmt.Printf("rows: %v\n", rows)
	fmt.Printf("ops: %v\n", ops)
}

func day1() {
	// rows, ops := readInputDay1("./input")
	rows, ops := readInputDay1("./input_test")
	// transpose
	numCols := len(rows[0])
	cols := make([][]int, numCols)

	for colIdx, _ := range rows[0] {
		cols[colIdx] = make([]int, len(rows))
		for rowIdx, _ := range rows {
			cols[colIdx][rowIdx] = rows[rowIdx][colIdx]
		}
	}

	sum := 0
	for i, col := range cols {
		result := col[0]

		mathOps := ops[i]
		for j := 1; j < len(col); j++ {
			switch mathOps {
			case ADD:
				result += col[j]
			case MULTIPLY:
				result *= col[j]
			}
		}
		sum += result
	}

	fmt.Println("part1: ", sum)
}

func readInputDay1(path string) ([][]int, []operators) {
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

func readInputDay2(path string) ([][]int, []operators) {
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
		fmt.Printf("fields: %v\n", fields)

	}
	return rows, ops
}
