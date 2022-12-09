package main

import (
    "aoc-2022/utils"
    "fmt"
    "os"
    "strconv"
)

type motion struct {
    dx int
    dy int
    repetitions int
}

type position struct {
    x int
    y int
}

func parseMotion(line string) motion {
    var dx int
    var dy int
    switch line[0] {
    case 'R':
        dx = 1
        dy = 0
    case 'L':
        dx = -1
        dy = 0
    case 'U':
        dx = 0
        dy = 1
    case 'D':
        dx = 0
        dy = -1
    default:
        panic(fmt.Sprintf("Unexpection motion direction %c", line[0]))
    }

    repetitions, err := strconv.Atoi(line[2:])
    if err != nil {
        panic(err)
    }

    return motion { dx, dy, repetitions }
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func move(head, tail *position, m motion, tailVisited map[position]struct{}) {
    for i := 0; i < m.repetitions; i++ {
        head.x += m.dx
        head.y += m.dy

        dx := head.x - tail.x
        dy := head.y - tail.y
        if abs(dx) > 1 || abs(dy) > 1 {
            if dx != 0 {
                tail.x += dx / abs(dx)
            }
            if dy != 0 {
                tail.y += dy / abs(dy)
            }
        }
        tailVisited[*tail] = struct{}{}
    }
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

    motions := utils.Map(lines, parseMotion)

    head := position { 0, 0 }
    tail := position { 0, 0 }
    tailVisited := make(map[position]struct{})
    for _, m := range motions {
        move(&head, &tail, m, tailVisited)
    }

    fmt.Printf("Answer: %d\n", len(tailVisited))
}
