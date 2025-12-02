package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	day1()
	// day2()
}

func day2() {
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// test()
	currentPosition := 50
	var correctTimes int
	for scanner.Scan() {
		line := scanner.Text()
		direction, movementStr := string(line[0]), line[1:]

		movement, err := strconv.Atoi(movementStr)
		if err != nil {
			panic(err)
		}
		var delta int = 1
		switch direction {
		case "L":
			delta = -1
		case "R":
		default:
			panic("unknown direction")
		}

		for _ = range movement {
			currentPosition += delta
			currentPosition = currentPosition % 100
			if currentPosition == 0 {
				correctTimes += 1
			}
		}
	}
	fmt.Println(correctTimes)

}

func day1() {
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// test()
	currentPosition := 50
	var correctTimes int
	for scanner.Scan() {
		line := scanner.Text()
		direction, movementStr := string(line[0]), line[1:]

		movement, err := strconv.Atoi(movementStr)
		if err != nil {
			panic(err)
		}
		var delta int = 1
		switch direction {
		case "L":
			delta = -1
		case "R":
		default:
			panic("unknown direction")
		}

		finalPos := ((movement * delta) + currentPosition) % 100
		if finalPos == 0 {
			correctTimes += 1
		}
		currentPosition = finalPos
	}
	fmt.Println(correctTimes)

}
