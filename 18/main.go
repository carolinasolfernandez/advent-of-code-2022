package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type cube struct {
	valid [][][]bool
	start [3]int
}

func main() {
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	cubes := map[int]map[int]map[int]bool{}

	for _, line := range lines {
		if line == "" {
			continue
		}
		sides := strings.Split(line, ",")
		x, err := strconv.Atoi(sides[0])
		if err != nil {
			panic("conversion")
		}
		y, err := strconv.Atoi(sides[1])
		if err != nil {
			panic("conversion")
		}
		z, err := strconv.Atoi(sides[2])
		if err != nil {
			panic("conversion")
		}

		if _, ok := cubes[x]; ok == false {
			cubes[x] = map[int]map[int]bool{}
		}
		if _, ok := cubes[x][y]; ok == false {
			cubes[x][y] = map[int]bool{}
		}
		cubes[x][y][z] = true
	}

	// Part 1
	sides := 0
	for x, yz := range cubes {
		for y, l := range yz {
			for z := range l {
				sides += 6
				for _, offset := range [][3]int{
					{-1, 0, 0}, {1, 0, 0}, {0, -1, 0}, {0, 1, 0}, {0, 0, -1}, {0, 0, 1},
				} {
					if _, ok := cubes[x+offset[0]][y+offset[1]][z+offset[2]]; ok {
						sides--
					}
				}
			}
		}
	}

	fmt.Println(sides)

	// Part 2
	sides = 0
	var xs, ys, zs []int
	for x, yz := range cubes {
		xs = append(xs, x)
		for y, l := range yz {
			ys = append(ys, y)
			for z := range l {
				zs = append(zs, z)
			}
		}
	}
	sort.Ints(xs)
	sort.Ints(ys)
	sort.Ints(zs)

	startX, startY, startZ := xs[0]-1, ys[0]-1, zs[0]-1

	c := cube{
		start: [3]int{startX, startY, startZ},
	}
	c.valid = make([][][]bool, xs[len(xs)-1]-startX+4)
	for x := range c.valid {
		c.valid[x] = make([][]bool, ys[len(ys)-1]-startY+4)
		for y := range c.valid[x] {
			c.valid[x][y] = make([]bool, zs[len(zs)-1]-startZ+4)
		}
	}
	sides = countSides(c, cubes, c.start[0], c.start[1], c.start[2])
	fmt.Println(sides)

}

func countSides(c cube, cubes map[int]map[int]map[int]bool, x, y, z int) int {
	c.valid[x-c.start[0]][y-c.start[1]][z-c.start[2]] = true
	sum := 0

	for _, offset := range [][3]int{
		{-1, 0, 0}, {1, 0, 0}, {0, -1, 0}, {0, 1, 0}, {0, 0, -1}, {0, 0, 1},
	} {
		cx, cy, cz := x+offset[0]-c.start[0], y+offset[1]-c.start[1], z+offset[2]-c.start[2]
		if cx < 0 || cx >= len(c.valid) || cy < 0 || cy >= len(c.valid[cx]) || cz < 0 || cz >= len(c.valid[cx][cy]) {
			continue
		} else if ok := c.valid[cx][cy][cz]; ok == true {
			continue
		}

		if cubes[x+offset[0]][y+offset[1]][z+offset[2]] {
			sum++
		} else {
			sum += countSides(c, cubes, x+offset[0], y+offset[1], z+offset[2])
		}
	}
	return sum
}
