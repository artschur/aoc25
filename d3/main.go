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
	p1 := 0
	for s.Scan() {
		bank := s.Text()
		p1 += getMaximumJoltage(bank, 2)
	}
	fmt.Println(p1)
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
			startSearchIndex = chosenChars[i-1].index + 1
		}
		limit := len(input) - digitsNeededAfter
		for j := startSearchIndex; j < limit; j++ {
			// if current is bigger than the last added into our arr
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
