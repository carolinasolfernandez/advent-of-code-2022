package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")

	start := 0
	end := 0
	cols := len(lines[0])
	rows := len(lines)
	grid := make([]rune, 0, rows*cols)
	valuesPart2 := make([][2]int, 0)

	for l, line := range lines {
		if line == "" {
			break
		}

		for k, char := range line {
			grid = append(grid, char)
			pos := l*cols + k
			if char == 'S' {
				start = pos
				grid[pos] = 'a'
				valuesPart2 = append(valuesPart2, [2]int{pos, 0})
			}
			if char == 'a' {
				valuesPart2 = append(valuesPart2, [2]int{pos, 0})
			}
			if char == 'E' {
				end = pos
				grid[pos] = 'z'
			}
		}
	}

	// Part 1
	gridPart1 := make([]rune, len(grid))
	copy(gridPart1, grid)

	valuesPart1 := [][2]int{{start, 0}}
	bestPart1 := getBest(valuesPart1, gridPart1, cols, rows, end)

	fmt.Println(bestPart1)

	// Part 2
	gridPart2 := make([]rune, len(grid))
	copy(gridPart2, grid)

	bestPart2 := getBest(valuesPart2, grid, cols, rows, end)

	fmt.Println(bestPart2)

}

func getBest(values [][2]int, grid []rune, cols, rows, end int) int {
	founds := make(map[int]int, 0)
	best := math.MaxInt

	for len(values) > 0 {
		current := values[0]
		values = values[1:]

		if v, ok := founds[current[0]]; ok && v < current[1] {
			continue
		}

		if current[1] > best {
			continue
		}

		if current[0] == end {
			best = current[1]
			continue
		}

		founds[current[0]] = founds[current[1]]
		x := current[0] % cols
		y := current[0] / cols

		for _, v := range []map[int]int{
			{x: y - 1},
			{x + 1: y},
			{x: y + 1},
			{x - 1: y},
		} {
			for nextX, nextY := range v {
				if nextX < 0 || nextY < 0 || nextX >= cols || nextY >= rows {
					continue
				}
				n := nextX + (cols * nextY)
				if grid[n] <= grid[current[0]]+1 {
					values = append(values, [2]int{n, current[1] + 1})
				}
			}
		}
	}
	return best
}
