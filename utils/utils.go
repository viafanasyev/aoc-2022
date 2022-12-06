package utils

import (
    "bufio"
    "os"
)

type Pair[T, U any] struct {
    First T
    Second U
}

func ReadLines(filePath string) ([]string, error) {
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

func SplitBy[T any](items []T, condition func(T) bool) [][]T {
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
    if len(currentChunk) > 0 {
        result = append(result, currentChunk)
    }
    return result
}

func Map[T, U any](items []T, f func(T) U) []U {
    result := make([]U, len(items))
    for i, item := range items {
        result[i] = f(item)
    }
    return result
}

func SumBy[T any](items []T, toInt func(T) int) int {
    sum := 0
    for _, item := range items {
        sum += toInt(item)
    }
    return sum
}

func Count[T any](items []T, condition func(T) bool) int {
    counter := 0
    for _, item := range items {
        if condition(item) {
            counter++
        }
    }
    return counter
}