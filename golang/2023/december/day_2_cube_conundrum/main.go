package main

import (
	"errors"
	"fmt"
	"strings"
    "strconv"
	"regexp"
)

type ColorId int
type BallCountByColor map[ColorId]int

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
    var match string
    var err error
    var temp int

    var r *regexp.Regexp = regexp.MustCompile("Game (\\d+)")
    match = r.FindAllStringSubmatch(string_splits[0], -1)[0][1]
    temp, err = strconv.Atoi(match)
    if err != nil {
        panic(err)
    }
    var game Game = Game{id: temp, sessions: []BallCountByColor{}}
    string_splits = strings.Split(string_splits[1], ";")


    r = regexp.MustCompile("(\\d+) (\\w+)")
    fmt.Println(r.FindAllStringSubmatch(string_splits[0], -1))
	return &game
}

func main() {
	_ = GameFromString("Game 97: 15 green, 9 blue; 14 blue, 14 red, 2 green; 18 red, 12 blue, 2 green")
}
