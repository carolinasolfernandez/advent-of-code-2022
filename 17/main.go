package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var (
	directions = map[byte]point{
		'<': {-1, 0},
		'>': {1, 0},
		'v': {0, -1},
	}

	rocks = []rock{
		// |..@@@@.|
		{{2, 0}, {3, 0}, {4, 0}, {5, 0}},

		// |...@...|
		// |..@@@..|
		// |...@...|
		{{3, 0}, {2, 1}, {3, 1}, {4, 1}, {3, 2}},

		// |....@..|
		// |....@..|
		// |..@@@..|
		{{2, 0}, {3, 0}, {4, 0}, {4, 1}, {4, 2}},

		// |..@....|
		// |..@....|
		// |..@....|
		// |..@....|
		{{2, 0}, {2, 1}, {2, 2}, {2, 3}},

		// |..@@...|
		// |..@@...|
		{{2, 0}, {3, 0}, {2, 1}, {3, 1}},
	}
)

type point struct{ x, y int }
type rock []point

func main() {
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	height := 0
	occ := map[point]bool{}
	cache := map[string]struct {
		round  int
		height int
	}{}
	contentIdx := 0

	for round := 1; round < 1000000000000; round++ {
		r := rocks[(round-1)%len(rocks)]

		newRock := make(rock, len(r))
		for i := 0; i < len(r); i++ {
			newRock[i].x = r[i].x
			newRock[i].y = r[i].y + height + 3
		}

		for {
			if contentIdx >= len(content) {
				contentIdx = 0
			}
			direction := content[contentIdx]
			contentIdx++

			var rockTemp rock
			for _, dir := range []point{directions[direction], directions['v']} {
				rockTemp = make(rock, len(newRock))
				for i, p := range newRock {
					p.x += dir.x
					p.y += dir.y

					if p.x < 0 || p.x > 6 || p.y < 0 {
						rockTemp = nil
						break
					}
					if _, ok := occ[p]; ok {
						rockTemp = nil
						break
					}
					rockTemp[i] = p
				}

				if rockTemp != nil {
					newRock = rockTemp
				}
			}

			if rockTemp != nil {
				continue
			}

			for _, p := range newRock {
				occ[p] = true
				if p.y+1 > height {
					height = p.y + 1
				}
			}

			if round == 2022 {
				fmt.Println("Part 1: ", height)
			}

			key := fmt.Sprintf("%d%d", (round-1)%len(rocks), contentIdx-1)
			if val, ok := cache[key]; ok {
				quotient := (1000000000000 - round) / (round - val.round)
				remainder := (1000000000000 - round) % (round - val.round)

				if remainder == 0 {
					fmt.Println("Part 2: ", height+(height-val.height)*quotient)
					os.Exit(0)
				}
			} else {
				cache[key] = struct {
					round  int
					height int
				}{round: round, height: height}
			}
			break
		}
	}
}
