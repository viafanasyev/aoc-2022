package main

import (
    "aoc-2022/utils"
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

    strategy := utils.Map(lines, stringToRoundStrategy)

    result := performStrategy(strategy)
    fmt.Printf("Answer: %d\n", result)
}
