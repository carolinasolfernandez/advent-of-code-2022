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

	stacks := map[int][]string{}

	stacks1Copy := map[int][]string{}
	stacks2Copy := map[int][]string{}

	set := true

	for _, line := range lines {
		if set {
			if string(line[1]) == "1" {
				set = false
				continue
			}
			stacks = setStack(stacks, line)
			continue
		}

		if line == "" {
			// Deep copy for part1 and part2
			for k, v := range stacks {
				stacks1Copy[k] = v
				stacks2Copy[k] = v
			}
			continue
		}

		words := strings.Split(line, " ")
		q := words[1]
		from := words[3]
		to := words[5]

		qq, _ := strconv.Atoi(q)
		fromq, _ := strconv.Atoi(from)
		toq, _ := strconv.Atoi(to)

		// Part 1
		var pop1 string
		qq1 := qq
		for qq1 > 0 {
			pop1, stacks1Copy[fromq] = stacks1Copy[fromq][0], stacks1Copy[fromq][1:]
			stacks1Copy[toq] = append([]string{pop1}, stacks1Copy[toq]...)
			qq1--
		}

		// Part 2
		var pop, pop2 []string
		pop, stacks2Copy[fromq] = stacks2Copy[fromq][0:qq], stacks2Copy[fromq][qq:]
		for _, v := range pop {
			pop2 = append(pop2, v)
		}
		stacks2Copy[toq] = append(pop2, stacks2Copy[toq]...)
	}
	fmt.Println(stacks1Copy[1][0], stacks1Copy[2][0], stacks1Copy[3][0], stacks1Copy[4][0], stacks1Copy[5][0], stacks1Copy[6][0], stacks1Copy[7][0], stacks1Copy[8][0], stacks1Copy[9][0])
	fmt.Println(stacks2Copy[1][0], stacks2Copy[2][0], stacks2Copy[3][0], stacks2Copy[4][0], stacks2Copy[5][0], stacks2Copy[6][0], stacks2Copy[7][0], stacks2Copy[8][0], stacks2Copy[9][0])
}

func setStack(stacks map[int][]string, line string) map[int][]string {
	pos := 1
	for col := 1; col <= 9; col++ {
		if col != 1 {
			pos += 4
		}
		if len(line) >= pos && string(line[pos]) != " " {
			stacks[col] = append(stacks[col], string(line[pos]))
		}
	}
	return stacks
}
