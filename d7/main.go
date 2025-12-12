package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	grid := readInput("./input")

	visited := make(map[string]bool)
	startingRow, startingCol := 0, 0
	for r, line := range grid {
		for c, char := range line {
			if char == "S" {
				startingRow = r
				startingCol = c
			}
		}
	}

	totalSplits := countSplits(startingRow, startingCol, &grid, visited)
	log.Printf("total splits: %d", totalSplits)
}

func countSplits(r, c int, grid *[][]string, visited map[string]bool) int {
	strCoord := fmt.Sprintf("%d,%d", r, c)
	limit := len(*grid) - 1
	if r == limit {
		return 0
	}
	if _, ok := visited[strCoord]; ok {
		return 0
	}
	totalSplits := 0

	visited[strCoord] = true
	if (*grid)[r][c] == "^" {
		totalSplits++
		totalSplits += countSplits(r+1, c+1, grid, visited)
		totalSplits += countSplits(r+1, c-1, grid, visited)
	} else {
		totalSplits += countSplits(r+1, c, grid, visited) // next row
	}

	return totalSplits
}

func readInput(path string) [][]string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	matrix := make([][]string, 0)
	for s.Scan() {
		matrixLine := make([]string, 0)
		line := strings.TrimSpace(s.Text())
		for _, char := range line {
			matrixLine = append(matrixLine, string(char))
		}
		matrix = append(matrix, matrixLine)
	}

	return matrix
}
