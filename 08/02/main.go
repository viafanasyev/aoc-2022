package main

import (
    "aoc-2022/utils"
    "fmt"
    "os"
)

const MAX_TREE_HEIGHT = 9

func updateScenicScore(treeGrid [][]int, scenicScoreGrid[][]int, x, y, dx, dy, maxX, maxY int) {
    currentlyVisibleTrees := make([]int, MAX_TREE_HEIGHT + 1)
    for 0 <= x && x < maxX && 0 <= y && y < maxY {
        currentTree := treeGrid[y][x]
        scenicScoreGrid[y][x] *= currentlyVisibleTrees[currentTree]
        for i := 0; i <= currentTree; i++ {
            currentlyVisibleTrees[i] = 1
        }
        for i := currentTree + 1; i < MAX_TREE_HEIGHT; i++ {
            currentlyVisibleTrees[i]++
        }
        x += dx
        y += dy
    }
}

func countScenicScore(treeGrid [][]int) [][]int {
    rowsNum := len(treeGrid)
    colsNum := len(treeGrid[0])
    scenicScoreGrid := make([][]int, rowsNum)
    for row := 0; row < rowsNum; row++ {
        scenicScoreGrid[row] = make([]int, colsNum)
        for col := 0; col < colsNum; col++ {
            scenicScoreGrid[row][col] = 1
        }
    }

    for row := 0; row < rowsNum; row++ {
        updateScenicScore(treeGrid, scenicScoreGrid, 0,           row, +1, 0, colsNum, rowsNum)
        updateScenicScore(treeGrid, scenicScoreGrid, colsNum - 1, row, -1, 0, colsNum, rowsNum)
    }
    for col := 0; col < colsNum; col++ {
        updateScenicScore(treeGrid, scenicScoreGrid, col, 0,           0, +1, colsNum, rowsNum)
        updateScenicScore(treeGrid, scenicScoreGrid, col, rowsNum - 1, 0, -1, colsNum, rowsNum)
    }

    return scenicScoreGrid
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

    scenicScoreGrid := countScenicScore(treeGrid)
    result := utils.MaxBy(scenicScoreGrid, utils.Max)
    fmt.Printf("Answer: %d\n", result)
}
