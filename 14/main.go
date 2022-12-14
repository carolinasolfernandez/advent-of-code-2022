package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")

	// Set Rocks
	rocks := map[float64]map[float64]int{}
	minX, _, maxX, maxY := setRocks(lines, rocks)

	// Deep Copy
	rocksPart1 := map[float64]map[float64]int{}
	rocksPart2 := map[float64]map[float64]int{}
	for k, v := range rocks {
		rocksPart1[k] = map[float64]int{}
		rocksPart2[k] = map[float64]int{}
		for k2, v2 := range v {
			rocksPart1[k][k2] = v2
			rocksPart2[k][k2] = v2
		}
	}

	// Part 1
	unitsPart1 := -1
	validPart1 := true
	for validPart1 {
		unitsPart1++
		validPart1 = simulateSandPart1(rocksPart1, minX, maxX, maxY, 500, 0)
	}
	fmt.Println(unitsPart1)

	// Part 2
	unitsPart2 := 0
	validPart2 := true
	for validPart2 {
		unitsPart2++
		validPart2 = simulateSandPart2(rocksPart2, minX, maxX, maxY, 500, 0)
	}
	fmt.Println(unitsPart2)

}

func setRocks(lines []string, rocks map[float64]map[float64]int) (minX, minY, maxX, maxY float64) {
	minX = math.MaxFloat64
	minY = math.MaxFloat64

	for _, line := range lines {
		if line == "" {
			break
		}
		pointsStr := regexp.MustCompile(" -> |,").Split(line, -1)
		var points []int
		for _, pointStr := range pointsStr {
			point, _ := strconv.Atoi(pointStr)
			points = append(points, point)
		}

		startX := float64(points[0])
		startY := float64(points[1])
		for i := 2; i < len(points); i += 2 {
			minX = math.Min(minX, startX)
			minY = math.Min(minY, startY)
			maxX = math.Max(maxX, startX)
			maxY = math.Max(maxY, startY)

			nextX := float64(points[i])
			nextY := float64(points[i+1])

			if startX != nextX {
				for j := math.Min(nextX, startX); j <= math.Max(nextX, startX); j++ {
					if _, ok := rocks[j]; ok == false {
						rocks[j] = map[float64]int{}
					}
					rocks[j][startY] = -1
				}
			} else {
				for j := math.Min(startY, nextY); j <= math.Max(startY, nextY); j++ {
					if _, ok := rocks[startX]; ok == false {
						rocks[startX] = map[float64]int{}
					}
					rocks[startX][j] = -1
				}
			}
			startX = nextX
			startY = nextY
		}
		minX = math.Min(minX, startX)
		minY = math.Min(minY, startY)
		maxX = math.Max(maxX, startX)
		maxY = math.Max(maxY, startY)
	}
	return
}

func simulateSandPart1(rocks map[float64]map[float64]int, minX, maxX, maxY, x, y float64) bool {
	if x == minX || x == maxX || y == maxY {
		return false
	}

	for _, move := range [][2]float64{{0, 1}, {-1, 1}, {1, 1}} {
		_, ok := rocks[x+move[0]][y+move[1]]
		if !ok {
			return simulateSandPart1(rocks, minX, maxX, maxY, x+move[0], y+move[1])
		}
	}

	if _, ok := rocks[x]; ok == false {
		rocks[x] = map[float64]int{}
	}
	rocks[x][y] = 1
	return true
}

func simulateSandPart2(rocks map[float64]map[float64]int, minX, maxX, maxY, x, y float64) bool {
	if y == maxY+1 {
		if _, ok := rocks[x]; ok == false {
			rocks[x] = map[float64]int{}
		}
		rocks[x][y] = 1
		return true
	}

	for _, move := range [][2]float64{{0, 1}, {-1, 1}, {1, 1}} {
		_, ok := rocks[x+move[0]][y+move[1]]
		if !ok {
			return simulateSandPart2(rocks, minX, maxX, maxY, x+move[0], y+move[1])
		}
	}

	if _, ok := rocks[x]; ok == false {
		rocks[x] = map[float64]int{}
	}
	rocks[x][y] = 1
	return x != 500 || y != 0
}
