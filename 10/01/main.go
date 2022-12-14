package main

import (
    "aoc-2022/utils"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type command struct {
    cycles int
    value int
}

func parseCommand(line string) command {
    splittedLine := strings.Fields(line)
    if splittedLine[0] == "noop" {
        return command { 1, 0 }
    } else if splittedLine[0] == "addx" {
        if len(splittedLine) < 2 {
            panic(fmt.Sprintf("Expected command value after 'addx' command, but got nothing"))
        }
        value, err := strconv.Atoi(splittedLine[1])
        if err != nil {
            panic(err)
        }
        return command { 2, value }
    } else {
        panic(fmt.Sprintf("Unexpected command '%s'", splittedLine[0]))
    }
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

    commands := utils.Map(lines, parseCommand)

    result := 0
    currentCycle := 1
    currentValue := 1
    nextInterestingCycle := 20
    for _, cmd := range commands {
        currentCycle += cmd.cycles
        if currentCycle > nextInterestingCycle {
            result += nextInterestingCycle * currentValue
            fmt.Printf("%d*%d = %d\n", nextInterestingCycle, currentValue, nextInterestingCycle * currentValue)
            nextInterestingCycle += 40
        }
        currentValue += cmd.value
    }
    fmt.Printf("Answer: %d\n", result)
}
