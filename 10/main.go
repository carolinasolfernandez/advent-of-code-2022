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

	value := 1
	cycles := []int{1}
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if line == "" {
			break
		}

		cmd := line[0:4]

		cycles = append(cycles, value)
		if cmd == "noop" {

		}
		if cmd == "addx" {
			val, _ := strconv.Atoi(line[5:])
			value += val
			cycles = append(cycles, value)
		}
	}

	// Part 1
	result := cycles[19]*20 + cycles[59]*60 + cycles[99]*100 + cycles[139]*140 + cycles[179]*180 + cycles[219]*220
	fmt.Println(result)

	// Part 2
	for cursor, sprite := range cycles {
		if cursor >= 200 {
			cursor = cursor - 200
		} else if cursor >= 160 {
			cursor = cursor - 160
		} else if cursor >= 120 {
			cursor = cursor - 120
		} else if cursor >= 80 {
			cursor = cursor - 80
		} else if cursor >= 40 {
			cursor = cursor - 40
		}

		if cursor+1 >= sprite && cursor+1 <= sprite+2 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		if (cursor+1)%40 == 0 {
			fmt.Print("\n")
		}
	}
}
