package main

import (
	"aoc2024/utils"
	"bufio"
	"fmt"
)

type Pos struct {
	Y, X int
}

func part1(scanner *bufio.Scanner) int {

	result := 0

	var matrix [][]rune

	for scanner.Scan() {
		matrix = append(matrix, []rune(scanner.Text()))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 'X' {

				result += check(matrix, i, j)
			}
		}
	}

	return result
}

func check(matrix [][]rune, v int, h int) int {
	height := len(matrix)
	if height == 0 {
		return 0
	}
	width := len(matrix[0])

	patterns := []map[Pos]rune{
		// Diagonal down-left
		{
			Pos{-1, -1}: 'M',
			Pos{-2, -2}: 'A',
			Pos{-3, -3}: 'S',
		},
		// Diagonal down-right
		{
			Pos{1, 1}: 'M',
			Pos{2, 2}: 'A',
			Pos{3, 3}: 'S',
		},
		// Diagonal up-right
		{
			Pos{-1, 1}: 'M',
			Pos{-2, 2}: 'A',
			Pos{-3, 3}: 'S',
		},
		// Diagonal up-left
		{
			Pos{1, -1}: 'M',
			Pos{2, -2}: 'A',
			Pos{3, -3}: 'S',
		},
		// Right horizontal
		{
			Pos{1, 0}: 'M',
			Pos{2, 0}: 'A',
			Pos{3, 0}: 'S',
		},
		// Left horizontal
		{
			Pos{-1, 0}: 'M',
			Pos{-2, 0}: 'A',
			Pos{-3, 0}: 'S',
		},
		// Down vertical
		{
			Pos{0, 1}: 'M',
			Pos{0, 2}: 'A',
			Pos{0, 3}: 'S',
		},
		// Up vertical
		{
			Pos{0, -1}: 'M',
			Pos{0, -2}: 'A',
			Pos{0, -3}: 'S',
		},
	}

	count := 0

	for _, pattern := range patterns {
		valid := true
		for pos := range pattern {
			newY := v + pos.X
			newX := h + pos.Y

			if newY < 0 || newY >= height || newX < 0 || newX >= width {
				valid = false
				break
			}
		}

		if valid {
			for pos, char := range pattern {
				if matrix[v+pos.X][h+pos.Y] != char {
					valid = false
					break
				}
			}
			if valid {
				count++
			}
		}
	}
	return count
}

func check2(matrix [][]rune, y int, x int) int {

	height := len(matrix)
	width := len(matrix[0])

	patterns := []struct {
		positions []Pos
		chars     []rune
	}{
		{
			// M.S
			// .A.
			// M.S
			positions: []Pos{{-1, -1}, {-1, 1}, {0, 0}, {1, -1}, {1, 1}},
			chars:     []rune{'M', 'S', 'A', 'M', 'S'},
		},
		{
			// S.M
			// .A.
			// S.M
			positions: []Pos{{-1, -1}, {-1, 1}, {0, 0}, {1, -1}, {1, 1}},
			chars:     []rune{'S', 'M', 'A', 'S', 'M'},
		},
		{
			// M.M
			// .A.
			// S.S
			positions: []Pos{{-1, -1}, {-1, 1}, {0, 0}, {1, -1}, {1, 1}},
			chars:     []rune{'M', 'M', 'A', 'S', 'S'},
		},
		{
			// S.S
			// .A.
			// M.M
			positions: []Pos{{-1, -1}, {-1, 1}, {0, 0}, {1, -1}, {1, 1}},
			chars:     []rune{'S', 'S', 'A', 'M', 'M'},
		},
	}

	count := 0

	for _, pattern := range patterns {
		valid := true
		for i, pos := range pattern.positions {
			newY := y + pos.Y
			newX := x + pos.X

			// Check bounds
			if newY < 0 || newY >= height || newX < 0 || newX >= width {
				valid = false
				break
			}

			// Check character match
			if matrix[newY][newX] != pattern.chars[i] {
				valid = false
				break
			}
		}

		if valid {
			count++
		}
	}

	return count
}

func part2(scanner *bufio.Scanner) int {
	result := 0

	var matrix [][]rune

	for scanner.Scan() {
		matrix = append(matrix, []rune(scanner.Text()))
	}

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[y]); x++ {

			if matrix[y][x] == 'A' {
				result += check2(matrix, y, x)
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
