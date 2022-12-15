package main

import (
    "aoc-2022/utils"
    "fmt"
    "math"
    "os"
    "strconv"
    "strings"
)

type monkey struct {
    items *utils.Queue[int]
    operation func(int) int
    divisor int
    monkeyIfTrue int
    monkeyIfFalse int
    itemsInspected int
}

func intPow(x, y int) int {
    return int(math.Pow(float64(x), float64(y)))
}

func parseOperation(line string) func(int) int {
    splitOp := strings.Fields(line)
    if len(splitOp) != 2 {
        panic(fmt.Sprintf("Expected opKind and opValue, but got '%s'", line))
    }
    opKind := splitOp[0]
    opValue, err := strconv.Atoi(splitOp[1])
    if err != nil {
        panic(err)
    }
    if opKind == "+" {
        return func(item int) int { return item + opValue }
    } else if opKind == "*" {
        return func(item int) int { return item * opValue }
    } else if opKind == "^" {
        return func(item int) int { return intPow(item, opValue) }
    } else {
        panic(fmt.Sprintf("Unexpected operation kind '%s'", opKind))
    }
}

func parseDivisor(line string) int {
    value, err := strconv.Atoi(line)
    if err != nil {
        panic(err)
    }

    return value
}

func parseMonkey(lines []string) monkey {
    if len(lines) != 5 {
        panic(fmt.Sprintf("Expected 5 lines for monkey, but got %d", len(lines)))
    }

    itemsArray := utils.Map(strings.Fields(lines[0]), func(x string) int {
        intX, err := strconv.Atoi(x)
        if err != nil {
            panic(err)
        }
        return intX
    })
    items := utils.NewQueue[int]()
    for _, item := range itemsArray {
        items.Enqueue(item)
    }

    operation := parseOperation(lines[1])

    divisor := parseDivisor(lines[2])

    monkeyIfTrue, err := strconv.Atoi(lines[3])
    if err != nil {
        panic(err)
    }

    monkeyIfFalse, err := strconv.Atoi(lines[4])
    if err != nil {
        panic(err)
    }

    return monkey { items, operation, divisor, monkeyIfTrue, monkeyIfFalse, 0 }
}

func makeMonkeyPlay(m *monkey, monkeys []monkey, modulo int) {
    for !m.items.IsEmpty() {
        item := m.items.Dequeue()
        item = m.operation(item)
        item = item % modulo
        if item % m.divisor == 0 {
            monkeys[m.monkeyIfTrue].items.Enqueue(item)
        } else {
            monkeys[m.monkeyIfFalse].items.Enqueue(item)
        }
        m.itemsInspected++
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

    monkeys := utils.Map(utils.SplitBy(lines, func(line string) bool { return len(line) == 0 }), parseMonkey)

    modulo := 1
    for _, m := range monkeys {
        modulo *= m.divisor
    }

    const rounds = 10000
    for round := 0; round < rounds; round++ {
        for i := 0; i < len(monkeys); i++ {
            makeMonkeyPlay(&monkeys[i], monkeys, modulo)
        }
    }

    mx1 := 0
    mx2 := 0
    for _, m := range monkeys {
        if m.itemsInspected > mx1 {
            mx2 = mx1
            mx1 = m.itemsInspected
        } else if m.itemsInspected > mx2 {
            mx2 = m.itemsInspected
        }
    }
    fmt.Printf("Max1: %d, Max2: %d\n", mx1, mx2)
    fmt.Printf("Answer: %d\n", mx1 * mx2)
}
