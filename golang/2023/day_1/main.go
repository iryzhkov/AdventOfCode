package main

import (
    "errors"
    "bufio"
    "fmt"
    "os"
    "log"
)


func readLines(file *os.File, lines chan string) {
    var scanner *bufio.Scanner = bufio.NewScanner(file)
    for scanner.Scan() {
        lines <- scanner.Text()
    }
    close(lines)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}


func extractNumber(line string) int {
    var first_number, second_number int = -1, 0
    for _, char := range(line) {
        if '0' <= char && char <= '9' {
            if first_number == -1 {
                first_number = second_number
            }
            second_number = int(char - '0')
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

    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var lines_channel chan string = make(chan string, 5)
    go readLines(file, lines_channel)

    var result int = 0
    for line := range lines_channel {
        result += extractNumber(line)
    }

    fmt.Println(result)
}
