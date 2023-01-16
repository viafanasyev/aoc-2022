package main

import (
    "aoc-2022/utils"
    "fmt"
    "os"
)

type listOrInt struct {
    intValue int
    listValue []listOrInt
}

func makeListNotInt(list []listOrInt) listOrInt {
    return listOrInt { -1, list }
}

func isListNotInt(it listOrInt) bool {
    return it.listValue != nil
}

func makeIntNotList(value int) listOrInt {
    return listOrInt { value, nil }
}

func isIntNotList(it listOrInt) bool {
    return it.listValue == nil
}

type packet struct {
    data listOrInt
}

func get(line string, pos *int, char byte) {
    if *pos >= len(line) {
        panic(fmt.Sprintf("Expected '%c', but got line end (line: '%s', pos: %d)", char, line, *pos))
    }
    if line[*pos] != char {
        panic(fmt.Sprintf("Expected '%c', but got '%c' (line: '%s', pos: %d)", char, line[*pos], line, *pos))
    }
    *pos = *pos + 1
}

func isDigit(char byte) bool {
    return char >= '0' && char <= '9'
}

func parseInt(line string, pos *int) listOrInt {
    value := 0
    for *pos < len(line) && isDigit(line[*pos]) {
        value = value * 10 + int(line[*pos] - '0')
        *pos = *pos + 1
    }
    return makeIntNotList(value)
}

func parseList(line string, pos *int) listOrInt {
    items := make([]listOrInt, 0)
    get(line, pos, '[')
    for {
        var item listOrInt
        if line[*pos] == '[' {
            item = parseList(line, pos)
        } else {
            item = parseInt(line, pos)
        }
        items = append(items, item)
        if line[*pos] == ']' {
            break
        }
        get(line, pos, ',')
    }
    get(line, pos, ']')
    return makeListNotInt(items)
}

func parsePacket(line string) packet {
    pos := 0
    return packet { parseList(line, &pos) }
}

func parsePacketPair(linePair []string) utils.Pair[packet, packet] {
    if len(linePair) != 2 {
        panic(fmt.Sprintf("Expected 2 lines in packet pair, but got %d", len(linePair)))
    }

    firstPacket := parsePacket(linePair[0])
    secondPacket := parsePacket(linePair[1])

    return utils.Pair[packet, packet] { firstPacket, secondPacket }
}

func compareListPair(firstList []listOrInt, secondList []listOrInt) int {
    for i := 0; i < len(firstList) && i < len(secondList); i++ {
        cmp := compareListOrIntPair(firstList[i], secondList[i])
        if cmp != 0 {
            return cmp
        }
    }
    return len(secondList) - len(firstList)
}

func compareListOrIntPair(firstListOrInt listOrInt, secondListOrInt listOrInt) int {
    if isIntNotList(firstListOrInt) && isIntNotList(secondListOrInt) {
        firstInt := firstListOrInt.intValue
        secondInt := secondListOrInt.intValue
        return secondInt - firstInt
    } else if isListNotInt(firstListOrInt) && isListNotInt(secondListOrInt) {
        firstList := firstListOrInt.listValue
        secondList := secondListOrInt.listValue
        return compareListPair(firstList, secondList)
    } else if isListNotInt(firstListOrInt) {
        firstList := firstListOrInt.listValue
        secondList := []listOrInt { secondListOrInt }
        return compareListPair(firstList, secondList)
    } else /* isListNotInt(secondListOrInt) */ {
        firstList := []listOrInt { firstListOrInt }
        secondList := secondListOrInt.listValue
        return compareListPair(firstList, secondList)
    }
}

func comparePacketPair(packetPair utils.Pair[packet, packet]) int {
    firstPacket := packetPair.First
    secondPacket := packetPair.Second
    return compareListOrIntPair(firstPacket.data, secondPacket.data)
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

    linePairs := utils.SplitBy(lines, func(line string) bool { return len(line) == 0 })
    packetPairs := utils.Map(linePairs, parsePacketPair)

    comparedPacketPairs := utils.Map(packetPairs, comparePacketPair)

    result := 0
    fmt.Printf("Ordered positions:")
    for i, cmp := range comparedPacketPairs {
        if cmp > 0 {
            result += i + 1
            fmt.Printf(" %d", i + 1)
        }
    }
    fmt.Printf("\n")
    fmt.Printf("Answer: %d\n", result)
}
