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

func findDuplicate(rucksackGroup []string) rune {
    if len(rucksackGroup) != 3 {
        panic(fmt.Sprintf("Expected 3 rucksacks, but got %d", len(rucksackGroup)))
    }

    for _, item := range rucksackGroup[0] {
        if strings.ContainsRune(rucksackGroup[1], item) && strings.ContainsRune(rucksackGroup[2], item) {
            return item
        }
    }
    panic(fmt.Sprintf("No intersection between %s, %s and %s is found", rucksackGroup[0], rucksackGroup[1], rucksackGroup[2]))
}

func sumBy[T any](items []T, toInt func(T) int) int {
    sum := 0
    for _, item := range items {
        sum += toInt(item)
    }
    return sum
}

func groupBy3[T any](items []T) [][]T {
    if len(items) % 3 != 0 {
        panic(fmt.Sprintf("Expecedd length to be divisible by 3, but got %d", len(items)))
    }

    groups := make([][]T, len(items) / 3)
    for i := 0; i < len(items); i += 3 {
        groups[i / 3] = []T{ items[i], items[i + 1], items[i + 2] }
    }
    return groups
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

    duplicates := mapTo(groupBy3(lines), findDuplicate)

    result := sumBy(duplicates, toPriority)
    fmt.Printf("Answer: %d\n", result)
}
