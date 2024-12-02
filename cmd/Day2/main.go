package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func scan() *bufio.Scanner {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	return scanner
}

func isSafe(levels []int) bool {
	if len(levels) < 2 {
		return true
	}

	increasing := levels[1] > levels[0]

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]

		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}

		if (increasing && diff < 0) || (!increasing && diff > 0) {
			return false
		}
	}

	return true
}

func part1(scanner *bufio.Scanner) int {

	result := 0

	for scanner.Scan() {
		splitted := strings.Fields(scanner.Text())
		ints := make([]int, len(splitted))

		for i, s := range splitted {
			num, err := strconv.Atoi(s)
			if err != nil {
				fmt.Printf("Error parsing number: %v\n", err)
				continue
			}
			ints[i] = num
		}

		if isSafe(ints) {
			result++
		}

	}
	return result
}

func part2(scanner *bufio.Scanner) int {

	result := 0

	for scanner.Scan() {
		splitted := strings.Fields(scanner.Text())
		ints := make([]int, len(splitted))

		for i, s := range splitted {
			num, err := strconv.Atoi(s)
			if err != nil {
				fmt.Printf("Error parsing number: %v\n", err)
				continue
			}
			ints[i] = num
		}

		safe := false

		if isSafe(ints) {
			safe = true
		} else {
			for i := range ints {
				excluded := make([]int, 0, len(ints)-1)
				excluded = append(excluded, ints[:i]...)
				excluded = append(excluded, ints[i+1:]...)

				if isSafe(excluded) {
					safe = true
					break
				}
			}
		}

		if safe {
			result++
		}

	}
	return result
}

func main() {
	scanner1 := scan()
	fmt.Println(part1(scanner1))

	scanner2 := scan()
	fmt.Println(part2(scanner2))
}
