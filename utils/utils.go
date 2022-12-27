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

func Sum(items []int) int {
    sum := 0
    for _, item := range items {
        sum += item
    }
    return sum
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

func CountTrue(items []bool) int {
    counter := 0
    for _, item := range items {
        if item {
            counter++
        }
    }
    return counter
}

func All[T any](items []T, condition func(T) bool) bool {
    for _, item := range items {
        if !condition(item) {
            return false
        }
    }
    return true
}

func None[T any](items []T, condition func(T) bool) bool {
    for _, item := range items {
        if condition(item) {
            return false
        }
    }
    return true
}

func Max(items []int) int {
    max := items[0]
    for _, item := range items[1:] {
        if item > max {
            max = item
        }
    }
    return max
}

func MaxBy[T any](items []T, toInt func(T) int) int {
    max := toInt(items[0])
    for _, item := range items[1:] {
        intItem := toInt(item)
        if intItem > max {
            max = intItem
        }
    }
    return max
}

func Min(items []int) int {
    min := items[0]
    for _, item := range items[1:] {
        if item < min {
            min = item
        }
    }
    return min
}

func MinBy[T any](items []T, toInt func(T) int) int {
    min := toInt(items[0])
    for _, item := range items[1:] {
        intItem := toInt(item)
        if intItem < min {
            min = intItem
        }
    }
    return min
}
