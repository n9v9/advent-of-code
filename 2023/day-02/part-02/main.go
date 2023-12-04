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

func (c *cubeCounter) countMax(color string, count int) {
	switch color {
	case "red":
		c.red = max(c.red, count)
	case "green":
		c.green = max(c.green, count)
	case "blue":
		c.blue = max(c.blue, count)
	}
}

func (c *cubeCounter) cubePower() int {
	return c.red * c.green * c.blue
}

func main() {
	var result int

	games := strings.FieldsFunc(input, func(r rune) bool { return r == '\n' })

	for _, game := range games {
		gameName, after, _ := strings.Cut(game, ": ")
		gameID, _ := strconv.Atoi(strings.Split(gameName, " ")[1])

		var cc cubeCounter

		for _, set := range strings.Split(after, "; ") {
			for _, cube := range strings.Split(set, ", ") {
				parts := strings.Split(cube, " ")
				count, _ := strconv.Atoi(parts[0])
				color := parts[1]
				cc.countMax(color, count)
			}
		}

		cubePower := cc.cubePower()
		fmt.Println("Game ID:", gameID)
		fmt.Println("Cube Powers:", cubePower)
		result += cubePower
	}

	fmt.Println("Result:", result)
}
