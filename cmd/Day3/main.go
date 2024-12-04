package main

import (
	"aoc2024/utils"
	"bufio"
	"fmt"
	"unicode"
)

func part1(scanner *bufio.Scanner) int {

	result := 0

	for scanner.Scan() {

		runes := []rune(scanner.Text())

		for i := 0; i < len(runes); i++ {
			if i+4 < len(runes) &&
				runes[i] == 'm' &&
				runes[i+1] == 'u' &&
				runes[i+2] == 'l' &&
				runes[i+3] == '(' {

				i += 4

				num1 := 0
				for i < len(runes) && unicode.IsDigit(runes[i]) {
					num1 = num1*10 + int(runes[i]-'0')
					i++
				}

				if i < len(runes) && runes[i] == ',' {
					i++

					num2 := 0
					for i < len(runes) && unicode.IsDigit(runes[i]) {
						num2 = num2*10 + int(runes[i]-'0')
						i++
					}

					if i < len(runes) && runes[i] == ')' {
						result += num1 * num2
					}
				}
			}
		}
	}
	return result
}

func part2(scanner *bufio.Scanner) int {
	result := 0

	enabled := true
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)

		for i := 0; i < len(runes); i++ {
			if i+6 < len(runes) &&
				runes[i] == 'd' &&
				runes[i+1] == 'o' &&
				runes[i+2] == 'n' &&
				runes[i+3] == '\'' &&
				runes[i+4] == 't' &&
				runes[i+5] == '(' &&
				runes[i+6] == ')' {
				enabled = false
				i += 6
				continue
			}

			if i+3 < len(runes) &&
				runes[i] == 'd' &&
				runes[i+1] == 'o' &&
				runes[i+2] == '(' &&
				runes[i+3] == ')' {
				enabled = true
				i += 3
				continue
			}

			if enabled && i+4 < len(runes) &&
				runes[i] == 'm' &&
				runes[i+1] == 'u' &&
				runes[i+2] == 'l' &&
				runes[i+3] == '(' {

				i += 4
				num1 := 0
				for i < len(runes) && unicode.IsDigit(runes[i]) {
					num1 = num1*10 + int(runes[i]-'0')
					i++
				}

				if i < len(runes) && runes[i] == ',' {
					i++
					num2 := 0
					for i < len(runes) && unicode.IsDigit(runes[i]) {
						num2 = num2*10 + int(runes[i]-'0')
						i++
					}

					if i < len(runes) && runes[i] == ')' {
						result += num1 * num2
					}
				}
			}
		}
	}
	return result
}

func main() {
	scanner1, cleanup, err := utils.OpenFileScanner("input.txt")

	if err != nil {
		fmt.Errorf("Scanner failed")
	}
	defer cleanup()

	result1 := part1(scanner1)
	fmt.Printf("Result part1: %d\n", result1)

	scanner2, cleanup, err := utils.OpenFileScanner("input.txt")

	if err != nil {
		fmt.Errorf("Scanner failed")
	}
	defer cleanup()

	result2 := part2(scanner2)
	fmt.Printf("Result part2: %d\n", result2)
}
