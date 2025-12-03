package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	p1, p2 := 0, 0
	for s.Scan() {
		bank := s.Text()
		p1 += getMaximumJoltage(bank, 2)  //part1
		p2 += getMaximumJoltage(bank, 12) //part2
	}
	fmt.Printf("p1: %v\n", p1)
	fmt.Printf("p2: %v\n", p2)
}

type joltageChar struct {
	value byte
	index int
}

func getMaximumJoltage(input string, lenght int) int {
	chosenChars := make([]joltageChar, lenght)
	for i := range lenght {
		digitsNeededAfter := lenght - i - 1
		startSearchIndex := 0
		if i > 0 {
			// start from last chosen char (rule points out after i choose a number i can only choose numbers after it)
			startSearchIndex = chosenChars[i-1].index + 1
		}
		limit := len(input) - digitsNeededAfter
		for j := startSearchIndex; j < limit; j++ {
			// find largest number in range
			if input[j] > chosenChars[i].value {
				chosenChars[i] = joltageChar{value: input[j], index: j}
			}
		}
	}

	total := 0
	for _, c := range chosenChars {
		digit := int(c.value - '0')
		total = total*10 + digit
	}
	return total
}
