package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	ranges, toCheck := readFile("./input")
	// fmt.Println(ranges)
	// fmt.Println(toCheck)
	acc := 0
	for _, num := range toCheck {
		inAnyRange := false
		for _, r := range ranges {
			if num >= r[0] && num <= r[1] {
				inAnyRange = true
				break
			}
		}
		if inAnyRange {
			acc++
		}
	}
	fmt.Println("Part 1:", acc)
}

func readFile(path string) (ranges [][]int, toCheck []int) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		// if we reached end of ranges...
		if line == "" {
			break
		}
		lowerBoundStr, upperBoundStr := strings.Split(line, "-")[0], strings.Split(line, "-")[1]
		lowerBound, err := strconv.Atoi(lowerBoundStr)
		if err != nil {
			panic(err)
		}
		upperBound, err := strconv.Atoi(upperBoundStr)
		if err != nil {
			panic(err)
		}
		ranges = append(ranges, []int{lowerBound, upperBound})
	}

	for s.Scan() {
		line := s.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		toCheck = append(toCheck, num)
	}

	return ranges, toCheck
}
