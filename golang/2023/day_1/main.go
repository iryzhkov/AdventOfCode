package main

import (
    "errors"
    "bufio"
    "fmt"
    "os"
    "log"
)


type trieNode struct {
    value int
    children map[rune]*trieNode
}


func addWord(root *trieNode, word *string, value int) {
    var currNode, tempNode *trieNode = root, nil
    for _, char := range(*word) {
        tempNode = currNode.children[char]
        if tempNode == nil {
            tempNode = &trieNode{value: -1, children: make(map[rune]*trieNode)}
            currNode.children[char] = tempNode
        }
        currNode = tempNode
    }
    currNode.value = value
}


func readLines(filename *string, lines chan string) {
    file, err := os.Open(*filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    var scanner *bufio.Scanner = bufio.NewScanner(file)
    for scanner.Scan() {
        lines <- scanner.Text()
    }
    close(lines)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}


func extractNumberPart1(line *string) int {
    var first_number, second_number int = -1, 0
    for _, char := range(*line) {
        if '0' <= char && char <= '9' {
            second_number = int(char - '0')
            if first_number == -1 {
                first_number = second_number
            }
        }
    }
    return first_number * 10 + second_number
}

func extractNumberPart2(line *string, trieRoot *trieNode) int {
    var first_number, second_number, temp_number int = -1, 0, -1
    var curr_nodes, temp_nodes *[]trieNode = &[]trieNode{*trieRoot}, nil
    var next_node *trieNode

    for _, char := range(*line) {
        temp_number = -1
        if '0' <= char && char <= '9' {
            temp_number = int(char - '0')
            curr_nodes = &[]trieNode{*trieRoot}
        } else if 'a' <= char && char <= 'z' {
            temp_nodes = &[]trieNode{*trieRoot}
            for _, node := range(*curr_nodes) {
                next_node = node.children[char]
                if next_node != nil {
                    *temp_nodes = append(*temp_nodes, *next_node)
                    if (*next_node).value != -1 {
                        temp_number = (*next_node).value
                    }
                }
            }
            curr_nodes = temp_nodes
        } else {
            curr_nodes = &[]trieNode{*trieRoot}
        }

        if temp_number != -1 {
            second_number = temp_number
            if first_number == -1 {
                first_number = second_number
            }
        }
    }
    return first_number * 10 + second_number
}


func main() {
    var filename string
    switch len(os.Args) {
        case 1:
            filename = "input.txt" 
        case 2:
            filename = os.Args[1]
        default:
            panic(errors.New("Only accept 1 or 2 arguments"))
    }

    var lines_channel chan string = make(chan string, 5)
    go readLines(&filename, lines_channel)

    var result int = 0
    for line := range lines_channel {
        result += extractNumberPart1(&line)
    }
    fmt.Printf("Part 1: %d\n", result)

    var number_strings [10]string = [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
    var trieRoot *trieNode = &trieNode{value: -1, children: make(map[rune]*trieNode)}
    for pos, number := range(number_strings) {
        addWord(trieRoot, &number, pos)
    }

    lines_channel = make(chan string, 5)
    go readLines(&filename, lines_channel)

    result = 0
    for line := range lines_channel {
        result += extractNumberPart2(&line, trieRoot)
    }
    fmt.Printf("Part 2: %d\n", result)
}
