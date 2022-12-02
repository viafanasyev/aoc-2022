package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
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

    caloriesChunks := splitBy(lines, func(line string) bool {
        return len(line) == 0
    })
    caloriesSums := mapTo(caloriesChunks, func(calories []string) int {
        sum := 0
        for _, calorie := range calories {
            calorieInt, err := strconv.Atoi(calorie)
            if err != nil {
                panic(err)
            }
            sum += calorieInt
        }
        return sum
    })

    sort.Slice(caloriesSums, func(i, j int) bool {
        return caloriesSums[i] > caloriesSums[j]
    })

    fmt.Printf("Top three sums: %d %d %d\n", caloriesSums[0], caloriesSums[1], caloriesSums[2])
    fmt.Printf("Answer: %d\n", caloriesSums[0] + caloriesSums[1] + caloriesSums[2])
}
