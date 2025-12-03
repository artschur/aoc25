package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("./input")
	// f, err := os.Open("./input_test")
	if err != nil {
		panic("failed opening file")
	}
	trimmed := strings.TrimSpace(string(content))
	p1 := 0
	p2 := 0
	for _, ids := range strings.Split(trimmed, ",") {
		bounds := strings.Split(ids, "-")
		lowerStr, upperStr := bounds[0], bounds[1]
		lower, err := strconv.Atoi(lowerStr)
		if err != nil {
			panic("failed converting lower bound to int")
		}
		upper, err := strconv.Atoi(upperStr)
		if err != nil {
			fmt.Println(upperStr, err)
			panic("failed converting upper bound to int")
		}

		for i := lower; i <= upper; i++ {
			stringedId := strconv.Itoa(i)
			if isInvalid(stringedId) {
				p1 += i
			}
			if isAtLeastTwice(stringedId) {
				p2 += i
			}
		}
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func isInvalid(id string) bool {
	n := len(id)

	if n%2 != 0 {
		return false
	}

	midpoint := n / 2
	left, right := id[:midpoint], id[midpoint:]

	return left == right
}

func isAtLeastTwice(id string) bool {
	if len(id) < 1 {
		return false
	}
	doubled := id + id
	doubled = doubled[1 : len(doubled)-1]

	return strings.Contains(doubled, id)
}
