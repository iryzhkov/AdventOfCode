package main

import (
    "errors"
    "bufio"
    "fmt"
    "os"
    "log"
)


func readLines(file *os.File, s chan string) {
    var scanner *bufio.Scanner = bufio.NewScanner(file)
    for scanner.Scan() {
        s <- scanner.Text()
    }
    close(s)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}


func extractNumber(line string) int {
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

    var s chan string = make(chan string, 5)
    go readLines(file, s)

    var result int = 0
    for line := range s {
        result += extractNumber(line)
    }

    fmt.Println(result)
}
