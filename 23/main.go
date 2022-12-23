package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

type elve struct {
	x         int
	y         int
	proposedX *int
	proposedY *int
}

var (
	moves = []string{"north", "south", "west", "east"}
	funcs = map[string]func(elve) (int, int){
		"north": func(i elve) (int, int) {
			return i.y - 1, i.x
		},
		"south": func(i elve) (int, int) {
			return i.y + 1, i.x
		},
		"west": func(i elve) (int, int) {
			return i.y, i.x - 1
		},
		"east": func(i elve) (int, int) {
			return i.y, i.x + 1
		},
	}
)

func main() {
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	elves := map[string]elve{}
	for i, line := range lines {
		if line == "" {
			break
		}

		for j, char := range line {
			if char == '#' {
				e := elve{
					y: i,
					x: j,
				}
				elves[fmt.Sprintf("%d-%d", i, j)] = e
			}
		}
	}

	move := 0
	elvesPart1 := map[string]elve{}
	for {
		// Propose
		for k, e := range elves {
			move0 := move % 4
			f := true

			if false == e.checkNorth(elves) && false == e.checkWest(elves) && false == e.checkSouth(elves) && false == e.checkEast(elves) {
				continue
			}
			for move0 != move%4 || f {
				f = false
				y, x := funcs[moves[move0]](e)
				var ok bool
				switch moves[move0] {
				case "north":
					ok = e.checkNorth(elves)
				case "south":
					ok = e.checkSouth(elves)
				case "east":
					ok = e.checkEast(elves)
				case "west":
					ok = e.checkWest(elves)
				}
				if ok == false {
					e.proposedX = &x
					e.proposedY = &y
					elves[k] = e
					break
				}
				move0++
				move0 = move0 % 4
			}
		}

		// check proposals
		for i, e1 := range elves {
			same := false
			if e1.proposedX == nil {
				continue
			}
			for j, e2 := range elves {
				if e2.proposedX == nil || i == j {
					continue
				}

				if *e1.proposedX == *e2.proposedX && *e1.proposedY == *e2.proposedY {
					e2.proposedX = nil
					e2.proposedY = nil
					elves[j] = e2
					same = true
				}
			}
			if same {
				e1.proposedX = nil
				e1.proposedY = nil
				elves[i] = e1
			}
		}

		elvesCopy := map[string]elve{}

		// execute proposal
		moved := false
		for _, e := range elves {
			if e.proposedX != nil {
				moved = true
				e.x = *e.proposedX
				e.y = *e.proposedY
			}
			e.proposedX = nil
			e.proposedY = nil
			if _, ok := elvesCopy[fmt.Sprintf("%d-%d", e.y, e.x)]; ok {
				panic("exists key")
			}
			elvesCopy[fmt.Sprintf("%d-%d", e.y, e.x)] = e
		}
		elves = elvesCopy
		if moved == false {
			break
		}
		if move == 9 {
			elvesPart1 = elvesCopy
		}
		move++
	}

	smallestX := math.MaxInt
	smallestY := math.MaxInt
	highestX := math.MinInt
	highestY := math.MinInt
	for _, e := range elvesPart1 {
		if e.x < smallestX {
			smallestX = e.x
		}
		if e.y < smallestY {
			smallestY = e.y
		}
		if e.x > highestX {
			highestX = e.x
		}
		if e.y > highestY {
			highestY = e.y
		}
	}

	fmt.Println("Part 1:", (highestX-smallestX+1)*(highestY-smallestY+1)-len(elvesPart1))
	fmt.Println("Part 2:", move+1)

}

func (e elve) checkNorth(elves map[string]elve) bool {
	row := e.y - 1
	col := e.x
	_, ok0 := elves[fmt.Sprintf("%d-%d", row, col)]
	_, ok1 := elves[fmt.Sprintf("%d-%d", row, col+1)]
	_, ok2 := elves[fmt.Sprintf("%d-%d", row, col-1)]
	return ok0 || ok1 || ok2
}

func (e elve) checkSouth(elves map[string]elve) bool {
	row := e.y + 1
	col := e.x
	_, ok0 := elves[fmt.Sprintf("%d-%d", row, col)]
	_, ok1 := elves[fmt.Sprintf("%d-%d", row, col+1)]
	_, ok2 := elves[fmt.Sprintf("%d-%d", row, col-1)]
	return ok0 || ok1 || ok2
}

func (e elve) checkEast(elves map[string]elve) bool {
	row := e.y
	col := e.x + 1
	_, ok0 := elves[fmt.Sprintf("%d-%d", row, col)]
	_, ok1 := elves[fmt.Sprintf("%d-%d", row+1, col)]
	_, ok2 := elves[fmt.Sprintf("%d-%d", row-1, col)]
	return ok0 || ok1 || ok2
}

func (e elve) checkWest(elves map[string]elve) bool {
	row := e.y
	col := e.x - 1
	_, ok0 := elves[fmt.Sprintf("%d-%d", row, col)]
	_, ok1 := elves[fmt.Sprintf("%d-%d", row+1, col)]
	_, ok2 := elves[fmt.Sprintf("%d-%d", row-1, col)]
	return ok0 || ok1 || ok2
}
