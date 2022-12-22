package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type encrypt struct {
	number int
	moved  bool
	pos    int
}

func main() {
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	var mixedPart1, mixedPart2 []encrypt
	for _, line := range lines {
		if line == "" {
			continue
		}

		number, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		encrypted := encrypt{
			number: number,
			pos:    len(mixedPart1),
		}
		mixedPart1 = append(mixedPart1, encrypted)

		encrypted.number = encrypted.number * 811589153
		mixedPart2 = append(mixedPart2, encrypted)
	}

	// Part 1
	for j := 0; j < len(mixedPart1); j++ {
		mixedPart1 = getMixed(mixedPart1, j)
	}

	// // Part 2
	for round := 0; round < 10; round++ {
		for j := 0; j < len(mixedPart2); j++ {
			mixedPart2 = getMixed(mixedPart2, j)
		}
	}

	// get start position
	var start1, start2 int
	for i := 0; i < len(mixedPart2); i++ {
		if mixedPart1[i].number == 0 {
			start1 = i
		}
		if mixedPart2[i].number == 0 {
			start2 = i
		}
	}

	var sum1, sum2 int
	for _, v := range []int{1000, 2000, 3000} {
		pos1 := v%len(mixedPart1) + start1
		if pos1 >= len(mixedPart1) {
			pos1 = pos1 - len(mixedPart1)
		}
		sum1 += mixedPart1[pos1].number

		pos2 := v%len(mixedPart2) + start2
		if pos2 >= len(mixedPart2) {
			pos2 = pos2 - len(mixedPart2)
		}
		sum2 += mixedPart2[pos2].number
	}
	fmt.Println("Part 1:", sum1)
	fmt.Println("Part 2:", sum2)

}

func getMixed(mixed []encrypt, j int) []encrypt {
	var i int
	for k2, v2 := range mixed {
		if v2.pos == j {
			i = k2
			break
		}
	}

	encrypted := mixed[i]
	mixed = append(mixed[:i], mixed[i+1:]...) // remove number from slice

	newIndex := i + encrypted.number%len(mixed)

	if encrypted.number > 0 && newIndex >= len(mixed) {
		newIndex = newIndex - len(mixed)
	} else if encrypted.number < 0 && newIndex <= 0 {
		newIndex = len(mixed) + newIndex
	}

	mixed = append(mixed[:newIndex+1], mixed[newIndex:]...) // create place for new element
	mixed[newIndex] = encrypted
	return mixed
}
