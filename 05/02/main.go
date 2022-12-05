package main

import (
    "aoc-2022/utils"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type crate rune

type command struct {
    from int
    to int
    count int
}

func parseStack(line string) *utils.Stack[crate] {
    crates := utils.NewStack[crate]()
    for _, cr := range line {
        crates.Push(crate(cr))
    }
    return crates
}

func parseCommand(line string) command {
    splittedLine := strings.Fields(line)
    if len(splittedLine) != 3 {
        panic(fmt.Sprintf("Expected 3 integers, splitted by spaces, but got %s", line))
    }

    count, err := strconv.Atoi(splittedLine[0])
    if err != nil {
        panic(err)
    }

    from, err := strconv.Atoi(splittedLine[1])
    if err != nil {
        panic(err)
    }

    to, err := strconv.Atoi(splittedLine[2])
    if err != nil {
        panic(err)
    }

    return command { from - 1, to - 1, count }
}

func interpret(cmd command, stacks []*utils.Stack[crate]) {
    tmpStack := utils.NewStack[crate]()
    for i := 0; i < cmd.count; i++ {
        value := stacks[cmd.from].Pop()
        tmpStack.Push(value)
    }
    for i := 0; i < cmd.count; i++ {
        value := tmpStack.Pop()
        stacks[cmd.to].Push(value)
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

    stacksAndCommands := utils.SplitBy(lines, func(line string) bool {
        return len(line) == 0
    })
    if len(stacksAndCommands) != 2 {
        fmt.Fprintf(
            os.Stderr,
            "Expected two groups of lines (stacks and commands), separated by space, but got %d group(s)\n",
            len(stacksAndCommands),
        )
        os.Exit(1)
    }

    stacks := utils.Map(stacksAndCommands[0], parseStack)
    commands := utils.Map(stacksAndCommands[1], parseCommand)

    for _, cmd := range commands {
        interpret(cmd, stacks)
    }

    fmt.Printf("Answer: ")
    for _, crates := range stacks {
        fmt.Printf("%c", crates.Peek())
    }
    fmt.Println()
}
