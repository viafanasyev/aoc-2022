package main

import (
    "aoc-2022/utils"
    "fmt"
    "math"
    "os"
)

type position struct {
    row int
    column int
}

func letterToHeight(letter rune) int {
    if letter >= 'a' && letter <= 'z' {
        return int(letter - 'a')
    } else if letter == 'S' {
        return letterToHeight('a')
    } else if letter == 'E' {
        return letterToHeight('z')
    } else {
        panic(fmt.Sprintf("Unexpected letter '%c'", letter))
    }
}

func linesToGrid(lines []string) ([][]int, []position, position) {
    grid := make([][]int, len(lines))
    starts := make([]position, 0)
    end := position { -1, -1 }
    for row, line := range lines {
        grid[row] = make([]int, len(line))
        for column, letter := range line {
            if letter == 'S' || letter == 'a' {
                starts = append(starts, position { row, column })
            } else if letter == 'E' {
                end = position { row, column }
            }
            grid[row][column] = letterToHeight(letter)
        }
    }

    if len(starts) == 0 {
        panic(fmt.Sprintf("Start cells not found"))
    }
    if end.row == -1 || end.column == -1 {
        panic(fmt.Sprintf("End cell not found"))
    }

    return grid, starts, end
}

func above(pos position) position {
    return position { pos.row - 1, pos.column }
}

func below(pos position) position {
    return position { pos.row + 1, pos.column }
}

func left(pos position) position {
    return position { pos.row, pos.column - 1 }
}

func right(pos position) position {
    return position { pos.row, pos.column + 1 }
}

func at(grid [][]int, pos position) int {
    return grid[pos.row][pos.column]
}

func legalPosition(grid [][]int, pos position) bool {
    if pos.row < 0 || pos.row >= len(grid) {
        return false
    }
    if pos.column < 0 || pos.column >= len(grid[pos.row]) {
        return false
    }
    return true
}

func stepAllowed(grid [][]int, from position, to position) bool {
    return legalPosition(grid, from) && legalPosition(grid, to) && at(grid, to) - at(grid, from) <= 1
}

func tryMakeStep(
    grid [][]int,
    currentPosition position,
    toPrevPos func(position) position,
    queue *utils.Queue[position],
    enqueued [][]bool,
    distance [][]int,
) {
    prevPos := toPrevPos(currentPosition)
    if stepAllowed(grid, prevPos, currentPosition) && !enqueued[prevPos.row][prevPos.column] {
        queue.Enqueue(prevPos)
        enqueued[prevPos.row][prevPos.column] = true
        distance[prevPos.row][prevPos.column] = at(distance, currentPosition) + 1
    }
}

func findShortestPaths(grid [][]int, to position) [][]int {
    enqueued := utils.Map(grid, func(row []int) []bool { return utils.Map(row, func(_ int) bool { return false }) })
    distance := utils.Map(grid, func(row []int) []int { return utils.Map(row, func(_ int) int { return math.MaxInt }) })
    queue := utils.NewQueue[position]()
    queue.Enqueue(to)
    enqueued[to.row][to.column] = true
    distance[to.row][to.column] = 0
    for !queue.IsEmpty() {
        currentPosition := queue.Dequeue()
        tryMakeStep(grid, currentPosition, above, queue, enqueued, distance)
        tryMakeStep(grid, currentPosition, below, queue, enqueued, distance)
        tryMakeStep(grid, currentPosition, left, queue, enqueued, distance)
        tryMakeStep(grid, currentPosition, right, queue, enqueued, distance)
    }
    return distance
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

    grid, starts, end := linesToGrid(lines)

    distance := findShortestPaths(grid, end)
    result := utils.MinBy(starts, func(start position) int { return at(distance, start) })
    fmt.Printf("Answer: %d\n", result)
}
