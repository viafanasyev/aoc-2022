package main

import (
    "aoc-2022/utils"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type fileNode struct {
    name string
    size int
    parent *fileNode
    children []*fileNode // nil, if it's a simple file. Non-nil splice, if it's a directory.
}

func newDir(name string) *fileNode {
    return &fileNode { name, 0, nil, make([]*fileNode, 0) };
}

func newFile(name string, size int) *fileNode {
    return &fileNode { name, size, nil, nil };
}

func (this *fileNode) pwd() string {
    if this.parent == nil {
        return this.name
    } else if this.parent.parent == nil && this.parent.name == "/" {
        // For case of '/' root dir
        return this.parent.name + this.name
    } else {
        return this.parent.pwd() + "/" + this.name
    }
}

func (this *fileNode) isDir() bool {
    return this.children != nil
}

func (this *fileNode) addChild(child *fileNode) {
    if child.parent != nil {
        panic(fmt.Sprintf(
            "File node %s already have a parent %s, but trying to add it to %s as a child",
            child.name,
            child.parent.name,
            this.name,
        ))
    }

    if !this.isDir() {
        panic(fmt.Sprintf("Trying to add child %s to file %s", child.name, this.name))
    }

    this.children = append(this.children, child)
    child.parent = this
}

func parseFileNodeFromLsEntry(line string) *fileNode {
    typeAndName := strings.Fields(line)
    if len(typeAndName) != 2 {
        panic(fmt.Sprintf("Expected type and name in file node entry, but got '%s'", line))
    }

    if typeAndName[0] == "dir" {
        return newDir(typeAndName[1])
    } else {
        size, err := strconv.Atoi(typeAndName[0])
        if err != nil {
            panic(err)
        }
        return newFile(typeAndName[1], size)
    }
}

func parseFileTree(terminalContents []string) *fileNode {
    root := newDir("/")
    curDir := root
    for _, terminalLine := range terminalContents {
        if terminalLine[0] == '$' {
            splittedCommand := strings.Fields(terminalLine[2:])
            if splittedCommand[0] == "cd" {
                dstDirName := splittedCommand[1]
                if dstDirName == ".." {
                    curDir = curDir.parent
                } else if dstDirName == "/" {
                    curDir = root
                } else {
                    var dstDir *fileNode = nil
                    for _, child := range curDir.children {
                        if child.name == dstDirName {
                            dstDir = child
                            break
                        }
                    }
                    if dstDir == nil {
                        dstDir = newDir(dstDirName)
                        curDir.addChild(dstDir)
                    }
                    curDir = dstDir
                }
            } else if splittedCommand[0] == "ls" {
                // Every non-dollar prefixed line is treated as 'ls' output, so just skip it
                continue
            } else {
                panic(fmt.Sprintf("Unknown command '%s'", splittedCommand[0]))
            }
        } else {
            child := parseFileNodeFromLsEntry(terminalLine)
            curDir.addChild(child)
        }
    }
    return root
}

func listContents(file *fileNode, depth int) {
    for i := 0; i < depth; i++ {
        fmt.Printf("  ")
    }

    if file.isDir() {
        fmt.Printf("- %s (dir)\n", file.name)
        for _, child := range file.children {
            listContents(child, depth + 1)
        }
    } else {
        fmt.Printf("- %s (file, size=%d)\n", file.name, file.size)
    }
}

func mapFilesToTotalSizes(root *fileNode) map[*fileNode]int {
    fileToTotalSize := make(map[*fileNode]int)

    var dfs func(*fileNode)
    dfs = func(curFile *fileNode) {
        fileToTotalSize[curFile] = curFile.size
        if curFile.isDir() {
            for _, child := range curFile.children {
                dfs(child)
                fileToTotalSize[curFile] += fileToTotalSize[child]
            }
        }
    }
    dfs(root)

    return fileToTotalSize
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

    root := parseFileTree(lines)

    //listContents(fileTree, 0)

    fileToTotalSize := mapFilesToTotalSizes(root)
    result := 0
    for file, totalSize := range fileToTotalSize {
        if file.isDir() && totalSize <= 100000 {
            result += totalSize
        }
    }
    fmt.Printf("Answer: %d\n", result)
}
