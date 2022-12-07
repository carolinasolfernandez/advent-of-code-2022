package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// A = rock = 1
// B = paper = 2
// C = scissor = 3
// X = rock/ lose = 0
// Y = paper/ draw = 3
// Z = scissor/ win = 6

func main() {
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	games := strings.Split(string(content), "\n")

	scoresPart1 := map[string]int{
		"AX": 1 + 3,
		"AY": 2 + 6,
		"AZ": 3 + 0,
		"BX": 1 + 0,
		"BY": 2 + 3,
		"BZ": 3 + 6,
		"CX": 1 + 6,
		"CY": 2 + 0,
		"CZ": 3 + 3,
	}
	scoresPart2 := map[string]int{
		"AX": 0 + 3,
		"AY": 3 + 1,
		"AZ": 6 + 2,
		"BX": 0 + 1,
		"BY": 3 + 2,
		"BZ": 6 + 3,
		"CX": 0 + 2,
		"CY": 3 + 3,
		"CZ": 6 + 1,
	}
	score1 := 0
	score2 := 0
	for _, game := range games {
		obj := strings.Split(game, " ")
		el := obj[0]
		yo := obj[1]
		score1 += scoresPart1[el+yo]
		score2 += scoresPart2[el+yo]
	}

	fmt.Println(score1)
	fmt.Println(score2)
}
