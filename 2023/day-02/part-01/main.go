package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type cubeCounter struct {
	red, green, blue int
}

func (c *cubeCounter) possible() bool {
	return c.red >= 0 && c.green >= 0 && c.blue >= 0
}

func (c *cubeCounter) subtract(color string, count int) {
	switch color {
	case "red":
		c.red -= count
	case "green":
		c.green -= count
	case "blue":
		c.blue -= count
	}
}

func main() {
	var result int

	games := strings.FieldsFunc(input, func(r rune) bool { return r == '\n' })

gameLoop:
	for _, game := range games {
		gameName, after, _ := strings.Cut(game, ": ")
		gameID, _ := strconv.Atoi(strings.Split(gameName, " ")[1])

		for _, set := range strings.Split(after, "; ") {
			cc := cubeCounter{
				red:   12,
				green: 13,
				blue:  14,
			}
			for _, cube := range strings.Split(set, ", ") {
				parts := strings.Split(cube, " ")
				count, _ := strconv.Atoi(parts[0])
				color := parts[1]
				cc.subtract(color, count)
			}
			if !cc.possible() {
				continue gameLoop
			}
		}

		fmt.Println("Game ID:", gameID)
		result += gameID
	}

	fmt.Println("Result:", result)
}
