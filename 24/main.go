package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	wall int = 1 << iota
	empty
	blizzardLeft
	blizzardRight
	blizzardDown
	blizzardUp
)

var valleyMap map[int]*valley

type valley struct {
	minute int
	valley [][]int
}

type node struct {
	pos pos
	val *valley
}

type pos struct {
	row, col int
}

func main() {
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")

	var grid [][]int
	var start, end pos

	for i, row := range lines {
		grid = append(grid, make([]int, len(row)))
		for j := range row {
			switch row[j] {
			case '#':
				grid[i][j] = wall
			case '.':
				grid[i][j] = empty
				if i == 0 {
					start = pos{row: 0, col: j}
				} else if i == len(lines)-1 {
					end = pos{row: len(grid) - 1, col: j}
				}
			case '<':
				grid[i][j] = blizzardLeft
			case '>':
				grid[i][j] = blizzardRight
			case '^':
				grid[i][j] = blizzardUp
			case 'v':
				grid[i][j] = blizzardDown
			default:
				panic("unrecognized char")
			}
		}
	}

	totalValleys := len(grid) * len(grid[0])
	valleyMap = make(map[int]*valley)

	valleyMap[0] = &valley{
		minute: 0,
		valley: grid,
	}

	for i := 1; i < totalValleys; i++ {
		newValley := make(map[pos]int)

		for rowNum, row := range grid {
			for colNum, tile := range row {
				p := pos{row: rowNum, col: colNum}
				if tile&wall != 0 {
					newValley[p] = wall
					continue
				}

				if tile&blizzardLeft != 0 {
					newValley[getOffset(grid, p, pos{row: 0, col: -1})] |= blizzardLeft
				}
				if tile&blizzardRight != 0 {
					newValley[getOffset(grid, p, pos{row: 0, col: 1})] |= blizzardRight
				}
				if tile&blizzardUp != 0 {
					newValley[getOffset(grid, p, pos{row: -1, col: 0})] |= blizzardUp
				}
				if tile&blizzardDown != 0 {
					newValley[getOffset(grid, p, pos{row: 1, col: 0})] |= blizzardDown
				}
			}
		}

		gridCopy := make([][]int, len(grid))
		for rowNum, row := range grid {
			gridCopy[rowNum] = make([]int, len(grid[rowNum]))
			for col := range row {
				p := pos{row: rowNum, col: col}
				if val, ok := newValley[p]; ok {
					gridCopy[rowNum][col] = val
				} else {
					gridCopy[rowNum][col] = empty
				}
			}
		}

		valleyMap[i] = &valley{
			minute: i,
			valley: gridCopy,
		}

		grid = gridCopy
	}

	startToGoal, ok := bFS(node{pos: start, val: valleyMap[0]}, end, totalValleys)
	if !ok {
		panic("no path")
	}

	fmt.Println("Part 1:", len(startToGoal)-1)

	// Part 2
	goalToStart, ok := bFS(startToGoal[len(startToGoal)-1], start, totalValleys)
	if !ok {
		panic("no path")
	}

	startToGoal2, ok := bFS(goalToStart[len(goalToStart)-1], end, totalValleys)
	if !ok {
		panic("no path")
	}

	fmt.Println("Part 2:", len(startToGoal)-1+len(goalToStart)-1+len(startToGoal2)-1)
}

func getOffset(grid [][]int, start, offset pos) pos {
	result := pos{row: start.row + offset.row, col: start.col + offset.col}

	if result.row == 0 {
		result.row = len(grid) - 2
	} else if result.row == len(grid)-1 {
		result.row = 1
	}

	if result.col == 0 {
		result.col = len(grid[result.row]) - 2
	} else if result.col == len(grid[result.row])-1 {
		result.col = 1
	}
	return result
}

func bFS(start node, end pos, totalValleys int) ([]node, bool) {
	var queue []node

	explored := map[node]bool{}
	parent := map[node]node{}

	queue = append(queue, start)

	for len(queue) != 0 {
		v := queue[0]
		queue = queue[1:]

		if v.pos == end {
			result := []node{v}

			if start == v {
				return result, true
			}

			curr := v
			for curr != start {
				curr = parent[curr]
				result = append(result, curr)
			}

			resultReversed := make([]node, len(result))
			for i, item := range result {
				resultReversed[len(resultReversed)-i-1] = item
			}
			return resultReversed, true
		}

		// Adjacent
		var nodes []node
		nextValleyMin := (v.val.minute + 1) % totalValleys
		nextValley := valleyMap[nextValleyMin]
		adjacent := adjacentPos(false, v.pos.row, v.pos.col, nextValley)
		for _, p := range adjacent {
			if nextValley.valley[p.row][p.col]&empty != 0 {
				n := node{
					pos: p,
					val: nextValley,
				}

				nodes = append(nodes, n)
			}
		}

		if nextValley.valley[v.pos.row][v.pos.col]&empty != 0 {
			n := node{
				pos: v.pos,
				val: nextValley,
			}

			nodes = append(nodes, n)
		}

		for _, neighbor := range nodes {
			if !explored[neighbor] {
				explored[neighbor] = true
				parent[neighbor] = v
				queue = append(queue, neighbor)
			}
		}
	}

	return nil, false
}

func adjacentPos(diag bool, row, col int, valley *valley) []pos {
	var results []pos

	for di := -1; di <= 1; di += 1 {
		for dj := -1; dj <= 1; dj += 1 {
			if di == 0 && dj == 0 {
				continue
			}

			absDj := dj
			absDi := di
			if di < 0 {
				absDi = -1 * di
			}
			if dj < 0 {
				absDj = -1 * dj
			}

			if !diag && absDi+absDj == 2 {
				continue
			}

			adjI := row + di
			adjJ := col + dj

			if !(adjI < 0 || adjJ < 0 || adjI >= len(valley.valley) || adjJ >= len(valley.valley[row])) {
				results = append(results, pos{row: adjI, col: adjJ})
			}
		}
	}

	return results
}
