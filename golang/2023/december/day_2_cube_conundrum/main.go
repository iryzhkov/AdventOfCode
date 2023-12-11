package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
    "bufio"
    "log"
    "os"
)

type ColorId int
type BallCountByColor map[ColorId]int

func (self BallCountByColor) isSubsetOf(other *BallCountByColor) bool {
	for color_id, self_num_balls := range self {
		other_num_balls := (*other)[color_id]
		if other_num_balls < self_num_balls {
			return false
		}
	}
	return true
}

func (self BallCountByColor) mergeWith(other *BallCountByColor) {
    for color_id, num_balls := range *other {
        curr_num_balls := self[color_id]
        self[color_id] = max(num_balls, curr_num_balls)
    }
}

func (self BallCountByColor) product() int {
    var result int = 1
    for _, num_balls := range self {
        result *= num_balls
    }
    return result
}

var color_list []string = []string{"blue", "green", "red"}

func stringToColorId(text string) (ColorId, error) {
	for pos, color := range color_list {
		if text == color {
			return ColorId(pos), nil
		}
	}
	return -1, errors.New("Color not found: " + text)
}

type Game struct {
	id       int
	sessions []BallCountByColor
}

func GameFromString(text string) *Game {
	var string_splits []string = strings.Split(text, ":")

	var r *regexp.Regexp = regexp.MustCompile("Game (\\d+)")
	match := r.FindAllStringSubmatch(string_splits[0], -1)[0][1]
	game_id, err := strconv.Atoi(match)
	if err != nil {
		panic(err)
	}

	var game Game = Game{id: game_id, sessions: []BallCountByColor{}}
	var session BallCountByColor
	r = regexp.MustCompile("(\\d+) (\\w+)")

	for _, split := range strings.Split(string_splits[1], ";") {
		session = BallCountByColor{}
		for _, matches := range r.FindAllStringSubmatch(split, -1) {
			num_balls, err := strconv.Atoi(matches[1])
			if err != nil {
				panic(err)
			}
			color_id, err := stringToColorId(matches[2])
			if err != nil {
				panic(err)
			}
			session[color_id] = num_balls
		}
		game.sessions = append(game.sessions, session)
	}

	return &game
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


func getFilename() string {
	var filename string
	switch len(os.Args) {
	case 1:
		filename = "input.txt"
	case 2:
		filename = os.Args[1]
	default:
		panic("Only accept 1 or 2 arguments")
	}
	return filename
}


func main() {
    var filename string = getFilename()
    var lines chan string = make(chan string, 5)
    var games []*Game = []*Game{}

    go readLines(&filename, lines)
    for line := range(lines) {
        games = append(games, GameFromString(line))
    }

    var target BallCountByColor = BallCountByColor{0: 14, 1: 13, 2: 12}
    var result int = 0
    for _, game := range games {
        var is_valid bool = true
        for _, session := range game.sessions {
            is_valid = is_valid && session.isSubsetOf(&target)
        }
        if is_valid {
            result += game.id
        }
    }
    fmt.Printf("Part 1 result: %v\n", result)

    result = 0
    for _, game := range games {
        var min_values BallCountByColor = BallCountByColor{0:0, 1:0, 2:0}
        for _, session := range game.sessions {
            min_values.mergeWith(&session)
        }
        result += min_values.product()
    }
    fmt.Printf("Part 2 result: %v\n", result)
}
