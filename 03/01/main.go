package main

import (
    "aoc-2022/utils"
    "fmt"
    "os"
    "strings"
)

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

    duplicates := utils.Map(lines, findDuplicate)

    result := utils.SumBy(duplicates, toPriority)
    fmt.Printf("Answer: %d\n", result)
}
