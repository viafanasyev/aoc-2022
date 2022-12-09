package main

import (
    "aoc-2022/utils"
    "fmt"
    "os"
)

func markVisibleTrees(treeGrid [][]int, visibleGrid[][]bool, x, y, dx, dy, maxX, maxY int) {
    highestTree := -1
    for 0 <= x && x < maxX && 0 <= y && y < maxY {
        if treeGrid[y][x] > highestTree {
            visibleGrid[y][x] = true
            highestTree = treeGrid[y][x]
        }
        x += dx
        y += dy
    }
}

func findVisibleTrees(treeGrid [][]int) [][]bool {
    rowsNum := len(treeGrid)
    colsNum := len(treeGrid[0])
    visibleGrid := make([][]bool, rowsNum)
    for row := 0; row < rowsNum; row++ {
        visibleGrid[row] = make([]bool, colsNum)
    }

    for row := 0; row < rowsNum; row++ {
        markVisibleTrees(treeGrid, visibleGrid, 0,           row, +1, 0, colsNum, rowsNum)
        markVisibleTrees(treeGrid, visibleGrid, colsNum - 1, row, -1, 0, colsNum, rowsNum)
    }
    for col := 0; col < colsNum; col++ {
        markVisibleTrees(treeGrid, visibleGrid, col, 0,           0, +1, colsNum, rowsNum)
        markVisibleTrees(treeGrid, visibleGrid, col, rowsNum - 1, 0, -1, colsNum, rowsNum)
    }

    return visibleGrid
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

    treeGrid := utils.Map(lines, func(line string) []int {
        treeLine := make([]int, len(line))
        for i := range line {
            treeLine[i] = int(line[i] - '0')
        }
        return treeLine
    })

    visibleTreeGrid := findVisibleTrees(treeGrid)
    result := utils.SumBy(visibleTreeGrid, utils.CountTrue)
    fmt.Printf("Answer: %d\n", result)
}
