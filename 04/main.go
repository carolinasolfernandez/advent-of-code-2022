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
	contains1 := 0
	contains2 := 0
	for _, line := range lines {
		if line == "" {
			break
		}
		elves := strings.Split(line, ",")
		sections1 := strings.Split(elves[0], "-")
		sections2 := strings.Split(elves[1], "-")
		sec11, _ := strconv.Atoi(sections1[0])
		sec12, _ := strconv.Atoi(sections1[1])
		sec21, _ := strconv.Atoi(sections2[0])
		sec22, _ := strconv.Atoi(sections2[1])
		// Part 1
		if (sec11 >= sec21 && sec12 <= sec22) || (sec21 >= sec11 && sec22 <= sec12) {
			contains1++
		}
		// Part 2
		if (sec11 >= sec21 && sec11 <= sec22) || (sec12 >= sec21 && sec12 <= sec22) || (sec22 >= sec12 && sec22 <= sec11) || (sec22 >= sec11 && sec22 <= sec12) {
			contains2++
		}
	}

	fmt.Println(contains1)
	fmt.Println(contains2)
}
