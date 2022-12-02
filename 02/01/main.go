package main

import (
    "bufio"
    "fmt"
    "os"
)

type shape int

const (
    ROCK shape = iota
    PAPER
    SCISSORS
)

const WIN_SCORE = 6
const DRAW_SCORE = 3
const LOSE_SCORE = 0

type roundStrategy struct {
    opponent shape
    player shape
}

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

func stringToRoundStrategy(line string) roundStrategy {
    if len(line) != 3 {
        panic(fmt.Sprintf("Unexpected round strategy %s", line))
    }

    var lhsShape shape = -1
    switch line[0] {
    case 'A':
        lhsShape = ROCK
    case 'B':
        lhsShape = PAPER
    case 'C':
        lhsShape = SCISSORS
    default:
        panic(fmt.Sprintf("Unexpected lhs in round strategy %s", line))
    }

    var rhsShape shape = -1
    switch line[2] {
    case 'X':
        rhsShape = ROCK
    case 'Y':
        rhsShape = PAPER
    case 'Z':
        rhsShape = SCISSORS
    default:
        panic(fmt.Sprintf("Unexpected rhs in round strategy %s", line))
    }

    if lhsShape == -1 || rhsShape == -1 {
        panic(fmt.Sprintf("Unexpected parsing result (%d, %d)", lhsShape, rhsShape))
    }

    return roundStrategy{ lhsShape, rhsShape }
}

func shapeToScore(sh shape) int {
    switch sh {
    case ROCK:
        return 1
    case PAPER:
        return 2
    case SCISSORS:
        return 3
    default:
        panic(fmt.Sprintf("Unexpected shape %d", sh))
    }
}

func outcomeScore(opponent, player shape) int {
   switch {
   case opponent == player:
       return DRAW_SCORE
   case opponent == ROCK && player == PAPER:
        return WIN_SCORE
   case opponent == ROCK && player == SCISSORS:
        return LOSE_SCORE
   case opponent == PAPER && player == ROCK:
        return LOSE_SCORE
   case opponent == PAPER && player == SCISSORS:
        return WIN_SCORE
   case opponent == SCISSORS && player == ROCK:
        return WIN_SCORE
   case opponent == SCISSORS && player == PAPER:
        return LOSE_SCORE
   default:
        panic(fmt.Sprintf("Unexpected shapes OP=%d, PL=%d", opponent, player))
   }
}

func performStrategy(strategy []roundStrategy) int {
    totalScore := 0
    for _, roundStrategy := range strategy {
        totalScore += shapeToScore(roundStrategy.player) + outcomeScore(roundStrategy.opponent, roundStrategy.player)
    }
    return totalScore
}

func main() {
    if len(os.Args) < 1 {
        fmt.Println("Expected input file")
        os.Exit(-1)
    }

    inputFilePath := os.Args[1]
    fmt.Printf("Input file: %s\n", inputFilePath)

    lines, err := readLines(inputFilePath)
    if err != nil {
        panic(err)
    }

    strategy := mapTo(lines, stringToRoundStrategy)

    result := performStrategy(strategy)
    fmt.Printf("Answer: %d\n", result)
}
