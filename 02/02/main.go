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

type outcome int

const (
    WIN outcome = iota
    DRAW
    LOSE
)

type roundStrategy struct {
    opponent shape
    player outcome
}

func stringToRoundStrategy(line string) roundStrategy {
    if len(line) != 3 {
        panic(fmt.Sprintf("Unexpected round strategy %s", line))
    }

    var opponent shape = -1
    switch line[0] {
    case 'A':
        opponent = ROCK
    case 'B':
        opponent = PAPER
    case 'C':
        opponent = SCISSORS
    default:
        panic(fmt.Sprintf("Unexpected lhs in round strategy %s", line))
    }

    var player outcome = -1
    switch line[2] {
    case 'X':
        player = LOSE
    case 'Y':
        player = DRAW
    case 'Z':
        player = WIN
    default:
        panic(fmt.Sprintf("Unexpected rhs in round strategy %s", line))
    }

    if opponent == -1 || player == -1 {
        panic(fmt.Sprintf("Unexpected parsing result (%d, %d)", opponent, player))
    }

    return roundStrategy{ opponent, player }
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

func outcomeToScore(out outcome) int {
    switch out {
    case WIN:
        return 6
    case DRAW:
        return 3
    case LOSE:
        return 0
    default:
        panic(fmt.Sprintf("Unexpected outcome %d", out))
    }
}

func toWinShape(sh shape) shape {
    switch sh {
    case ROCK:
        return PAPER
    case PAPER:
        return SCISSORS
    case SCISSORS:
        return ROCK
    default:
        panic(fmt.Sprintf("Unexpected shape %d", sh))
    }
}

func toLoseShape(sh shape) shape {
    switch sh {
    case ROCK:
        return SCISSORS
    case PAPER:
        return ROCK
    case SCISSORS:
        return PAPER
    default:
        panic(fmt.Sprintf("Unexpected shape %d", sh))
    }
}

func roundScore(opponent shape, player outcome) int {
    switch player {
    case WIN:
        return shapeToScore(toWinShape(opponent)) + outcomeToScore(player)
    case DRAW:
        return shapeToScore(opponent) + outcomeToScore(player)
    case LOSE:
        return shapeToScore(toLoseShape(opponent)) + outcomeToScore(player)
    default:
        panic(fmt.Sprintf("Unexpected outcome %d", player))
    }
}

func performStrategy(strategy []roundStrategy) int {
    totalScore := 0
    for _, roundStrategy := range strategy {
        totalScore += roundScore(roundStrategy.opponent, roundStrategy.player)
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
