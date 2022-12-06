package main

import (
    "aoc-2022/utils"
    "fmt"
    "os"
)

func zeroOrOne(x int) bool {
    return x == 0 || x == 1
}

func findFirstGroupOfUnique(line string, groupLength int) int {
    if len(line) < groupLength {
        return -1
    }

    numOfOccurances := make([]int, int('z') - int('a') + 1)
    for i := 0; i < groupLength; i++ {
        numOfOccurances[line[i] - 'a']++
    }

    if utils.All(numOfOccurances, zeroOrOne) {
        return 0
    }

    for pos := 0; pos + groupLength < len(line); pos++ {
        numOfOccurances[line[pos] - 'a']--
        numOfOccurances[line[pos + groupLength] - 'a']++
        if utils.All(numOfOccurances, zeroOrOne) {
            return pos + 1
        }
    }
    return -1
}

func main() {
    if len(os.Args) < 2 {
        fmt.Fprintf(os.Stderr, "Expected input file\n")
        os.Exit(1)
    }

    inputFilePath := os.Args[1]
    fmt.Printf("Input file: %s\n", inputFilePath)

    lines, err := utils.ReadLines(inputFilePath)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error reading file: %s\n", err)
        os.Exit(1)
    }

    if len(lines) != 1 {
        fmt.Fprintf(os.Stderr, "Expected single line, but got %d\n", len(lines))
        os.Exit(1)
    }

    line := lines[0]
    const groupLength = 4
    pos := findFirstGroupOfUnique(line, groupLength)
    if pos == -1 {
        fmt.Fprintf(os.Stderr, "No group of unique characters of length %d is found in '%s'\n", groupLength, line)
        os.Exit(1)
    }
    fmt.Printf("Answer: %d\n", pos + groupLength)
}
