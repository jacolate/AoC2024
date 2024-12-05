package main

import (
	"aoc2024/utils"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Rule struct {
	num, after int
}

type Update struct {
	position map[int]int 
	length   int         
}

func parse(scanner *bufio.Scanner) ([]Rule, []Update) {
	var rules []Rule
	var updates []Update

	trigger := true

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			trigger = false
			continue
		}

		if trigger {
			parsed := strings.Split(line, "|")

			num, err1 := strconv.Atoi(parsed[0])
			after, err2 := strconv.Atoi(parsed[1])
			if err1 == nil && err2 == nil {
				rules = append(rules, Rule{num: num, after: after})
			}
		} else {
			parsed := strings.Split(line, ",")

			position := make(map[int]int)

			for i, val := range parsed {
				num, err := strconv.Atoi(val)
				if err != nil {
					fmt.Errorf("Parsing failed")
				}
				position[num] = i
			}
			updates = append(updates, Update{position: position, length: len(position)})
		}
	}
	return rules, updates
}

func part1(scanner *bufio.Scanner) int {
	result := 0

	rules, updates := parse(scanner)

	for _, update := range updates {
		if isValid(update, rules) {
			result += middle(update)
		}
	}

	return result
}

func part2(scanner *bufio.Scanner) int {
	result := 0

	rules, updates := parse(scanner)

	for _, update := range updates {
		if !isValid(update, rules) {
			reorder(update, rules) 
			result += middle(update)
		}
	}

	return result
}

func isValid(update Update, rules []Rule) bool {
	for _, rule := range rules {
		posNum, existsNum := update.position[rule.num]
		posAfter, existsAfter := update.position[rule.after]

		if existsNum && existsAfter && posNum >= posAfter {
			return false
		}
	}
	return true
}

func reorder(update Update, rules []Rule) {
	for {
		changed := false
		for _, rule := range rules {
			posNum, existsNum := update.position[rule.num]
			posAfter, existsAfter := update.position[rule.after]

			if existsNum && existsAfter && posNum < posAfter {
				update.position[rule.num], update.position[rule.after] = update.position[rule.after], update.position[rule.num]
				changed = true
			}
		}

		if !changed {
			break
		}
	}
}

func middle(update Update) int {
	pages := make([]int, len(update.position))
	for page, pos := range update.position {
		pages[pos] = page
	}
	return pages[len(pages)/2]
}

func main() {
	scanner1, cleanup1, err := utils.OpenFileScanner("input.txt")
	if err != nil {
		fmt.Errorf("Scanner failed")
		return
	}
	defer cleanup1()

	result1 := part1(scanner1)
	fmt.Printf("Result part1: %d\n", result1)

	scanner2, cleanup2, err := utils.OpenFileScanner("input.txt")
	if err != nil {
		fmt.Errorf("Scanner failed")
		return
	}
	defer cleanup2()

	result2 := part2(scanner2)
	fmt.Printf("Result part2: %d\n", result2)
}
