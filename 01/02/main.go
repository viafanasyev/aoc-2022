package main

import (
    "aoc-2022/utils"
    "fmt"
    "os"
    "sort"
    "strconv"
)

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

    caloriesChunks := utils.SplitBy(lines, func(line string) bool {
        return len(line) == 0
    })
    caloriesSums := utils.Map(caloriesChunks, func(calories []string) int {
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
