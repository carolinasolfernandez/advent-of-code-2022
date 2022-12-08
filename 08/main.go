package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	var treesHorizontal [][]int
	treesVertical := make([][]int, len(lines[0]))

	// Set Variables with file
	for k, line := range lines {
		var treeLine []int
		for i := 0; i < len(line); i++ {
			if len(treesVertical[i]) == 0 {
				treesVertical[i] = make([]int, len(lines))
			}

			height, _ := strconv.Atoi(string(line[i]))
			treeLine = append(treeLine, height)
			treesVertical[i][k] = height
		}
		treesHorizontal = append(treesHorizontal, treeLine)
	}

	visibleHorizontal := deepCopyMultiInt(treesHorizontal)

	// Part 1
	for i := 0; i < len(treesHorizontal); i++ {
		for j := 0; j < len(treesHorizontal[i]); j++ {
			lineH := treesHorizontal[i]
			lineV := treesVertical[j]
			r := checkRight(j, lineH)
			if r == 0 {
				r = checkLeft(j, lineH)
			}
			if r == 0 {
				r = checkRight(i, lineV)
			}
			if r == 0 {
				r = checkLeft(i, lineV)
			}
			visibleHorizontal[i][j] = r
		}
	}

	quantity := 0
	for h, treeL := range visibleHorizontal {
		for v, tree := range treeL {
			if tree > 0 || h == 0 || v == 0 || h == len(visibleHorizontal)-1 || v == len(treeL)-1 {
				quantity++
			}

		}
	}

	fmt.Println(quantity)

	// Part 2
	scoreMatrix := deepCopyMultiInt(treesHorizontal)
	for i := 0; i < len(treesHorizontal); i++ {
		for j := 0; j < len(treesHorizontal[i]); j++ {
			lineH := treesHorizontal[i]
			lineV := treesVertical[j]
			right := getScoreRight(j, lineH)
			left := getScoreLeft(j, lineH)
			top := getScoreRight(i, lineV)
			bottom := getScoreLeft(i, lineV)
			scoreMatrix[i][j] = right * left * top * bottom
		}
	}

	maxScore := 0
	for _, scoreL := range scoreMatrix {
		for _, score := range scoreL {
			if score > maxScore {
				maxScore = score
			}

		}
	}
	fmt.Println(maxScore)

}

func checkLeft(pos int, line []int) int {
	height := line[pos]
	for l := 0; l < pos; l++ {
		if height <= line[l] {
			return 0
		}
	}
	return height
}

func checkRight(pos int, line []int) int {
	height := line[pos]
	for l := pos + 1; l < len(line); l++ {
		if height <= line[l] {
			return 0
		}
	}
	return height
}

func getScoreLeft(pos int, trees []int) int {
	score := 0
	for l := pos - 1; l >= 0; l-- {
		score++
		if trees[l] >= trees[pos] {
			break
		}
	}
	return score
}

func getScoreRight(pos int, trees []int) int {
	score := 0
	for l := pos + 1; l < len(trees); l++ {
		score++
		if trees[l] >= trees[pos] {
			break
		}
	}
	return score
}

func deepCopyMultiInt(original [][]int) [][]int {
	deepCopy := make([][]int, len(original))
	for i := range original {
		deepCopy[i] = make([]int, len(original[i]))
		copy(deepCopy[i], original[i])
	}
	return deepCopy
}
