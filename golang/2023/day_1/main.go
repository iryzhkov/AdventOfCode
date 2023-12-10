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


func addWord(root_node *trieNode, word string, value int) {
    var curr_node, temp_node *trieNode = root_node, nil
    for _, char := range(word) {
        temp_node = curr_node.children[char]
        if temp_node == nil {
            temp_node = &trieNode{value: -1, children: make(map[rune]*trieNode)}
            curr_node.children[char] = temp_node
        }
        curr_node = temp_node
    }
    curr_node.value = value
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


func extractNumberPart1(line string) int {
    var first_number, second_number int = -1, 0
    for _, char := range(line) {
        if '0' <= char && char <= '9' {
            second_number = int(char - '0')
            if first_number == -1 {
                first_number = second_number
            }
        }
    }
    return first_number * 10 + second_number
}


func extractNumberPart2(line string, root_node *trieNode) int {
    const unset int = -1
    var first_number, second_number, temp_number int = unset, 0, unset 
    var curr_nodes, temp_nodes *[]trieNode = &[]trieNode{*root_node}, nil
    var next_node *trieNode

    for _, char := range(line) {
        temp_number = unset 

        if '0' <= char && char <= '9' {

            temp_number = int(char - '0')
            curr_nodes = &[]trieNode{*root_node}

        } else if 'a' <= char && char <= 'z' {

            temp_nodes = &[]trieNode{*root_node}

            for _, node := range(*curr_nodes) {
                next_node = node.children[char]
                if next_node != nil {
                    *temp_nodes = append(*temp_nodes, *next_node)
                    if (*next_node).value != unset {
                        temp_number = (*next_node).value
                    }
                }
            }

            curr_nodes = temp_nodes

        } else {

            curr_nodes = &[]trieNode{*root_node}

        }

        if temp_number != unset {
            second_number = temp_number
            if first_number == unset {
                first_number = second_number
            }
        }
    }
    return first_number * 10 + second_number
}


func getFilename() string {
    var filename string
    switch len(os.Args) {
        case 1:
            filename = "input.txt" 
        case 2:
            filename = os.Args[1]
        default:
            panic(errors.New("Only accept 1 or 2 arguments"))
    }
    return filename
}


func part1(filename string) {
    var lines_channel chan string = make(chan string, 5)
    go readLines(&filename, lines_channel)

    var result int = 0
    for line := range lines_channel {
        result += extractNumberPart1(line)
    }
    fmt.Printf("Part 1: %d\n", result)
}


func part2(filename string) {
    var number_strings [10]string = [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
    var trieRoot trieNode = trieNode{value: -1, children: make(map[rune]*trieNode)}
    for pos, number := range(number_strings) {
        addWord(&trieRoot, number, pos)
    }

    var lines_channel chan string = make(chan string, 5)
    go readLines(&filename, lines_channel)

    var result int = 0
    for line := range lines_channel {
        result += extractNumberPart2(line, &trieRoot)
    }
    fmt.Printf("Part 2: %d\n", result)
}


func main() {
    var filename string = getFilename()
    part1(filename)
    part2(filename)
}
