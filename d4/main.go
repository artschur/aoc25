package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// day1()
	day2()
}

func day2() {
	lines := readInput("./input_test")

	maxAdjacent := 4
	totalRemoved := 0
	for {
		removedInIteration := 0
		for line := range lines {
			for col, _ := range lines[line] {
				if lines[line][col] == "@" && validatePosition(lines, line, col, maxAdjacent) {
					removePosition(lines, line, col)
					removedInIteration++
				}
			}
		}
		if removedInIteration == 0 {
			break
		}
		totalRemoved += removedInIteration
	}
	fmt.Println("Part 2:", totalRemoved)
}

func day1() {
	lines := readInput("./input")

	maxAdjacent := 4
	p1 := 0
	for line := range lines {
		for col, _ := range lines[line] {
			if lines[line][col] == "@" && validatePosition(lines, line, col, maxAdjacent) {
				p1++
			}
		}
	}
	fmt.Println("Part 1:", p1)
}

func removePosition(lines [][]string, row, col int) {
	lines[row][col] = "."
}

func validatePosition(lines [][]string, row, col, maxAdjacent int) bool {
	positionsToCheck := [8][2]int{
		//top
		{row - 1, col - 1}, {row - 1, col}, {row - 1, col + 1},
		//middle
		{row, col - 1}, {row, col + 1},
		//bottom
		{row + 1, col - 1}, {row + 1, col}, {row + 1, col + 1},
	}

	count := 0
	for _, pos := range positionsToCheck {
		newRow, newCol := pos[0], pos[1]
		if newRow < 0 || newRow >= len(lines) || newCol < 0 || newCol >= len(lines[newRow]) {
			continue
		}
		if lines[newRow][newCol] == "@" {
			count++
			if count >= maxAdjacent {
				return false
			}
		}
	}
	return true
}

func readInput(path string) [][]string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var lines [][]string
	for s.Scan() {
		line := s.Text()
		arr := []string{}
		for i := 0; i < len(line); i++ {
			arr = append(arr, string(line[i]))
		}
		lines = append(lines, arr)
	}
	return lines
}
