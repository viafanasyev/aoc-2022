package main

import (
    "aoc-2022/utils"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type interval struct {
    from int
    to int
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

func lineToIntervalPair(line string) utils.Pair[interval, interval] {
    strPair := strings.FieldsFunc(line, func(c rune) bool { return c == ',' })
    if len(strPair) != 2 {
        panic(fmt.Sprintf("Expected single comma in line, but got %s", line))
    }

    firstInterval := parseInterval(strPair[0])
    secondInterval := parseInterval(strPair[1])
    return utils.Pair[interval, interval] { firstInterval, secondInterval }
}

func fullyContains(intervalPair utils.Pair[interval, interval]) bool {
    firstInterval := intervalPair.First
    secondInterval := intervalPair.Second
    if firstInterval.from <= secondInterval.from && secondInterval.to <= firstInterval.to {
        return true
    }
    if secondInterval.from <= firstInterval.from && firstInterval.to <= secondInterval.to {
        return true
    }
    return false
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

    intervalPairs := utils.Map(lines, lineToIntervalPair)
    result := utils.Count(intervalPairs, fullyContains)
    fmt.Printf("Answer: %d\n", result)
}
