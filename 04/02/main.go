package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type interval struct {
    from int
    to int
}

type pair[T, U any] struct {
    first T
    second U
}

func readLines(filePath string) ([]string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

func splitBy[T any](items []T, condition func(T) bool) [][]T {
    result := [][]T{}
    currentChunk := []T{}
    for _, item := range items {
        if condition(item) {
            if len(currentChunk) > 0 {
                result = append(result, currentChunk)
            }
            currentChunk = make([]T, 0)
        } else {
            currentChunk = append(currentChunk, item)
        }
    }
    return result
}

func mapTo[T, U any](items []T, f func(T) U) []U {
    result := make([]U, len(items))
    for i, item := range items {
        result[i] = f(item)
    }
    return result
}

func sumBy[T any](items []T, toInt func(T) int) int {
    sum := 0
    for _, item := range items {
        sum += toInt(item)
    }
    return sum
}

func count[T any](items []T, condition func(T) bool) int {
    counter := 0
    for _, item := range items {
        if condition(item) {
            counter++
        }
    }
    return counter
}

func parseInterval(str string) interval {
    strInterval := strings.FieldsFunc(str, func(c rune) bool { return c == '-' })
    if len(strInterval) != 2 {
        panic(fmt.Sprintf("Expected single '-' in interval, but got %s", str))
    }

    from, err := strconv.Atoi(strInterval[0])
    if err != nil {
        panic(err)
    }

    to, err := strconv.Atoi(strInterval[1])
    if err != nil {
        panic(err)
    }

    if from > to {
        panic(fmt.Sprintf("Left bound of interval mustn't be greater than right bound, but got [%d; %d]", from, to))
    }

    return interval { from, to }
}

func lineToIntervalPair(line string) pair[interval, interval] {
    strPair := strings.FieldsFunc(line, func(c rune) bool { return c == ',' })
    if len(strPair) != 2 {
        panic(fmt.Sprintf("Expected single comma in line, but got %s", line))
    }

    firstInterval := parseInterval(strPair[0])
    secondInterval := parseInterval(strPair[1])
    return pair[interval, interval] { firstInterval, secondInterval }
}

func overlap(intervalPair pair[interval, interval]) bool {
    firstInterval := intervalPair.first
    secondInterval := intervalPair.second
    if firstInterval.to < secondInterval.from {
        return false
    }
    if secondInterval.to < firstInterval.from {
        return false
    }
    return true
}

func main() {
    if len(os.Args) < 2 {
        fmt.Fprintf(os.Stderr, "Expected input file\n")
        os.Exit(1)
    }

    inputFilePath := os.Args[1]
    fmt.Printf("Input file: %s\n", inputFilePath)

    lines, err := readLines(inputFilePath)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error reading file: %s\n", err)
        os.Exit(1)
    }

    intervalPairs := mapTo(lines, lineToIntervalPair)
    result := count(intervalPairs, overlap)
    fmt.Printf("Answer: %d\n", result)
}
