package main

import (
    "bufio"
    "fmt"
    "sort"
    "strconv"
    "strings"
    "aoc2024/utils"
)

func part1(scanner *bufio.Scanner) int {
    var l1, l2 []int

    for scanner.Scan() {
        splitted := strings.Fields(scanner.Text())
        n1, _ := strconv.Atoi(splitted[0])
        n2, _ := strconv.Atoi(splitted[1])
        
        l1 = append(l1, n1)
        l2 = append(l2, n2)
    }

    sort.Ints(l1)
    sort.Ints(l2)
    var result int

    for i, val := range l1 {
        if val-l2[i] < 0 {
            result += l2[i] - val
        } else {
            result += val - l2[i]
        }
    }

    return result
}

func part2(scanner *bufio.Scanner) int {
    var l1, l2 []int

    for scanner.Scan() {
        splitted := strings.Fields(scanner.Text())
        n1, _ := strconv.Atoi(splitted[0])
        n2, _ := strconv.Atoi(splitted[1])
        
        l1 = append(l1, n1)
        l2 = append(l2, n2)
    }

    var result int
    counts := make(map[int]int)

    for _, val := range l2 {
        counts[val]++
    }

    for _, val := range l1 {
        result += val * counts[val]
    }

    return result
}

func main() {
    scanner, cleanup, err := utils.OpenFileScanner("input.txt")
    if err != nil {
        panic(err)
    }
    defer cleanup()
    fmt.Printf("Result Part 1: %d\n", part1(scanner))
    
    scanner, cleanup, err = utils.OpenFileScanner("input.txt")
    if err != nil {
        panic(err)
    }
    defer cleanup()
    fmt.Printf("Result Part 2: %d\n", part2(scanner))
}
