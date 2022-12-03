package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

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

func toPriority(item rune) int {
    if item >= 'a' && item <= 'z' {
        return int(item) - 'a' + 1
    } else if item >= 'A' && item <= 'Z' {
        return int(item) - 'A' + 27
    } else {
        panic(fmt.Sprintf("Unexpected item %c", item))
    }
}

func findDuplicate(rucksackItems string) rune {
    itemCount := len(rucksackItems)
    if itemCount % 2 != 0 {
        panic(fmt.Sprintf("Expected compartments to be equal in length, but got %s", rucksackItems))
    }

    leftCompartmentItems := rucksackItems[:(itemCount / 2)]
    rightCompartmentItems := rucksackItems[(itemCount / 2):]

    for _, leftCompartmentItem := range leftCompartmentItems {
        if strings.ContainsRune(rightCompartmentItems, leftCompartmentItem) {
            return leftCompartmentItem
        }
    }
    panic(fmt.Sprintf("No intersection between %s and %s is found", leftCompartmentItems, rightCompartmentItems))
}

func sumBy[T any](items []T, toInt func(T) int) int {
    sum := 0
    for _, item := range items {
        sum += toInt(item)
    }
    return sum
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

    duplicates := mapTo(lines, findDuplicate)

    result := sumBy(duplicates, toPriority)
    fmt.Printf("Answer: %d\n", result)
}
