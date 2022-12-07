package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	scores := map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10,
		"k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20,
		"u": 21, "v": 22, "w": 23, "x": 24, "y": 25, "z": 26,
		"A": 27, "B": 28, "C": 29, "D": 30, "E": 31, "F": 32, "G": 33, "H": 34, "I": 35, "J": 36,
		"K": 37, "L": 38, "M": 39, "N": 40, "O": 41, "P": 42, "Q": 43, "R": 44, "S": 45, "T": 46,
		"U": 47, "V": 48, "W": 49, "X": 50, "Y": 51, "Z": 52,
	}
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	/*
		// Part 1
			score := 0
			for _, line := range lines {
				if line == "" {
					break
				}
				itemsCount := len(line)
				first := line[:itemsCount/2]
				sec := line[itemsCount/2:]
				char := getCoincidence(itemsCount/2, first, sec)
				score = score + scores[char]
			}
	*/

	// Part 2
	score := 0
	i := 0
	for {
		line1 := lines[i]
		if line1 == "" {
			break
		}
		i++
		line2 := lines[i]
		i++
		line3 := lines[i]
		i++
		char := getCoincidence2(line1, line2, line3)
		score = score + scores[char]

	}
	fmt.Println(score)
}

func getCoincidence(itemsCount int, first, sec string) string {
	for i := 0; i < itemsCount; i++ {
		char := string(first[i])
		fmt.Println(char)
		for j := 0; j < itemsCount; j++ {
			charS := string(sec[j])
			if charS == char {
				return char
			}
		}
	}
	panic(errors.New("no coincidence"))
}
func getCoincidence2(first, sec, third string) string {
	for i := 0; i < len(first); i++ {
		char := string(first[i])
		for j := 0; j < len(sec); j++ {
			charS := string(sec[j])
			for k := 0; k < len(third); k++ {
				charT := string(third[k])
				if (charT == char) && (charT == charS) {
					return char
				}
			}
		}
	}
	panic(errors.New("no coincidence"))
}
