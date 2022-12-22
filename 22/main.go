package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	var matrix [][]int
	var moves string
	for j, line := range lines {
		if line == "" {
			moves = lines[j+1]
			break
		}

		var row []int
		for i := 0; i < len(line); i++ {
			var val int
			if string(line[i]) == " " {
				val = -1
			} else if string(line[i]) == "." {
				val = 1
			} else if string(line[i]) == "#" {
				val = 0
			}
			row = append(row, val)
		}
		matrix = append(matrix, row)
	}

	k := 0
	row := 0
	col := 0
	facing := 0
	for k < len(moves) {
		if moves[k] == 'R' {
			facing = (facing + 1) % 4
			k++
		} else if moves[k] == 'L' {
			facing--
			if facing < 0 {
				facing = 3
			}
			k++
		}

		j := k
		for k < len(moves) && moves[k] != 'R' && moves[k] != 'L' {
			k++
		}

		move, err := strconv.Atoi(moves[j:k])
		if err != nil {
			panic("conversion error")
		}
		lastCol := col
		lastRow := row

		for move > 0 {
			if facing == 0 { // right
				col++
				if col >= len(matrix[row]) {
					col = 0
				}
			} else if facing == 1 { // up
				row++
				if row >= len(matrix) {
					row = 0
				}
				if col >= len(matrix[row]) {
					continue
				}
			} else if facing == 2 { // right
				col--
				if col < 0 {
					col = len(matrix[row]) - 1
				}
			} else if facing == 3 { // down
				row--
				if row < 0 {
					row = len(matrix) - 1
				}
				if col >= len(matrix[row]) {
					continue
				}
			}
			val := matrix[row][col]
			if val == -1 {
				continue
			}
			if val == 0 {
				row = lastRow
				col = lastCol
				break
			}
			if val == 1 {
				lastRow = row
				lastCol = col
			}
			move--
		}

	}
	fmt.Println(1000*(row+1) + 4*(col+1) + facing)
}
