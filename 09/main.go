package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	positionsCountPart1 := map[int]map[int]int{}
	positionsCountPart2 := map[int]map[int]int{}
	currentHX := 0
	currentHY := 0
	currentTX := 0
	currentTY := 0
	knotsX := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	knotsY := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, line := range lines {
		if line == "" {
			break
		}
		position, _ := strconv.Atoi(line[2:])
		for position > 0 {
			position--
			if string(line[0]) == "R" {
				currentHX += 1
			}
			if string(line[0]) == "L" {
				currentHX -= 1
			}
			if string(line[0]) == "U" {
				currentHY += 1
			}
			if string(line[0]) == "D" {
				currentHY -= 1
			}

			// Part 1
			currentTY, currentTX = compare(currentTY, currentTX, currentHY, currentHX)
			if _, ok := positionsCountPart1[currentTX]; ok {
				positionsCountPart1[currentTX][currentTY] = 1
			} else {
				positionsCountPart1[currentTX] = map[int]int{currentTY: 1}
			}

			// Part 2
			knotsX[0] = currentTX
			knotsY[0] = currentTY
			for i := 1; i < 9; i++ {
				knotsY[i], knotsX[i] = compare(knotsY[i], knotsX[i], knotsY[i-1], knotsX[i-1])
			}

			if _, ok := positionsCountPart2[knotsX[8]]; ok {
				positionsCountPart2[knotsX[8]][knotsY[8]] = 1
			} else {
				positionsCountPart2[knotsX[8]] = map[int]int{knotsY[8]: 1}
			}
		}
	}

	positionsPart1 := 0
	for _, v := range positionsCountPart1 {
		for range v {
			positionsPart1++
		}
	}
	fmt.Println(positionsPart1)

	positionsPart2 := 0
	for _, v := range positionsCountPart2 {
		for range v {
			positionsPart2++
		}
	}
	fmt.Println(positionsPart2)
}

func compare(tailY, tailX, headY, headX int) (int, int) {
	if tailY == headY && (math.Abs(float64(headX)-float64(tailX))) > 1 {
		if (headX - tailX) > 0 {
			tailX += 1
		} else {
			tailX -= 1
		}
	} else if headX == tailX && (math.Abs(float64(headY)-float64(tailY))) > 1 {
		if (headY - tailY) > 0 {
			tailY += 1
		} else {
			tailY -= 1
		}
	} else if (math.Abs(float64(headY)-float64(tailY))) > 1 || (math.Abs(float64(headX)-float64(tailX))) > 1 {
		if (headY - tailY) > 0 {
			tailY += 1
		} else {
			tailY -= 1
		}
		if (headX - tailX) > 0 {
			tailX += 1
		} else {
			tailX -= 1
		}
	}
	return tailY, tailX
}
