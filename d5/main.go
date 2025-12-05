package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// day1()
	day2()
}

func day2() {
	ranges, _ := readFile("./input", true)

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	var merged [][]int
	if len(ranges) > 0 {
		merged = append(merged, ranges[0])
	}

	for i := 1; i < len(ranges); i++ {
		current := ranges[i]
		lastMergedIndex := len(merged) - 1
		previous := merged[lastMergedIndex]

		// 4-10 and 8-12
		if current[0] <= previous[1] {
			if current[1] > previous[1] {
				merged[lastMergedIndex][1] = current[1]
			}
		} else {
			merged = append(merged, current)
		}
	}

	sum := 0
	for _, r := range merged {
		sum += r[1] - r[0] + 1
	}

	fmt.Println("Part 2:", sum)
}

func day1() {
	ranges, toCheck := readFile("./input", false)
	// fmt.Println(ranges)
	// fmt.Println(toCheck)
	sum := 0
	for _, num := range toCheck {
		inAnyRange := false
		for _, r := range ranges {
			if num >= r[0] && num <= r[1] {
				inAnyRange = true
				break
			}
		}
		if inAnyRange {
			sum++
		}
	}
	fmt.Println("Part 1:", sum)
}

func readFile(path string, dayTwo bool) (ranges [][]int, toCheck []int) {
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

	if dayTwo {
		return ranges, nil
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
