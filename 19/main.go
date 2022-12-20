package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type blueprint struct {
	id                int
	oreRobotCost      int
	clayRobotCost     int
	obsidianRobotCost [2]int
	geodeRobotCost    [2]int
}

func main() {
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	resultPart1 := 0
	resultPart2 := 1
	for i, line := range lines {
		if line == "" {
			continue
		}
		blue := blueprint{}
		_, err := fmt.Sscanf(line,
			"Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&blue.id, &blue.oreRobotCost, &blue.clayRobotCost, &blue.obsidianRobotCost[0], &blue.obsidianRobotCost[1], &blue.geodeRobotCost[0], &blue.geodeRobotCost[1],
		)
		if err != nil {
			panic("scan error")
		}
		fmt.Println("Calculating Blueprint", blue.id)

		cache := map[string]int{}
		resultPart1 += blue.id * getMaxGeodes(blue, 1, 0, 0, 0, 0, 0, 0, 0, 24, cache, 0)

		if i < 3 {
			resultPart2 *= getMaxGeodes(blue, 1, 0, 0, 0, 0, 0, 0, 0, 32, cache, 0)
		}
	}

	fmt.Println("Part 1: ", resultPart1)
	fmt.Println("Part 2: ", resultPart2)
}

func getMaxGeodes(b blueprint, oreRobots, clayRobots, obsidianRobots, geodeRobots, ore, clay, obsidian, geode, time int, cache map[string]int, currentMax int) int {
	if time <= 0 {
		return geode
	}

	cacheKey := fmt.Sprintf("%d-%d-%d-%d-%d-%d-%d-%d-%d", oreRobots, clayRobots, obsidianRobots, geodeRobots, ore, clay, obsidian, geode, time)
	if v, e := cache[cacheKey]; e {
		return v
	}

	potentialMax := geode + geodeRobots*time + time*(time+1)/2
	if currentMax >= potentialMax {
		cache[cacheKey] = 0
		return 0
	}

	oreCosts := []int{b.oreRobotCost, b.clayRobotCost, b.obsidianRobotCost[0], b.geodeRobotCost[0]}
	sort.Ints(oreCosts)
	if oreRobots > oreCosts[len(oreCosts)-1] || clayRobots >= b.obsidianRobotCost[1] || obsidianRobots >= b.geodeRobotCost[1] {
		cache[cacheKey] = 0
		return 0
	}

	max := geode

	// Create geodes robot
	if b.geodeRobotCost[0] <= ore && b.geodeRobotCost[1] <= obsidian {
		result := getMaxGeodes(b, oreRobots, clayRobots, obsidianRobots, geodeRobots+1, ore+oreRobots-b.geodeRobotCost[0], clay+clayRobots, obsidian+obsidianRobots-b.geodeRobotCost[1], geode+geodeRobots, time-1, cache, max)
		if result > max {
			max = result
		}
		cache[cacheKey] = max

		return max
	}

	// Create obsidian robot
	if b.obsidianRobotCost[0] <= ore && b.obsidianRobotCost[1] <= clay {
		result := getMaxGeodes(b, oreRobots, clayRobots, obsidianRobots+1, geodeRobots, ore+oreRobots-b.obsidianRobotCost[0], clay+clayRobots-b.obsidianRobotCost[1], obsidian+obsidianRobots, geode+geodeRobots, time-1, cache, max)
		if result > max {
			max = result
		}
	}

	// Create clay robot
	if b.clayRobotCost <= ore {
		result := getMaxGeodes(b, oreRobots, clayRobots+1, obsidianRobots, geodeRobots, ore+oreRobots-b.clayRobotCost, clay+clayRobots, obsidian+obsidianRobots, geode+geodeRobots, time-1, cache, max)
		if result > max {
			max = result
		}
	}

	// Create ore robot
	if b.oreRobotCost <= ore {
		result := getMaxGeodes(b, oreRobots+1, clayRobots, obsidianRobots, geodeRobots, ore+oreRobots-b.oreRobotCost, clay+clayRobots, obsidian+obsidianRobots, geode+geodeRobots, time-1, cache, max)
		if result > max {
			max = result
		}
	}

	result := getMaxGeodes(b, oreRobots, clayRobots, obsidianRobots, geodeRobots, ore+oreRobots, clay+clayRobots, obsidian+obsidianRobots, geode+geodeRobots, time-1, cache, max)
	if result > max {
		max = result
	}

	cache[cacheKey] = max

	return max

}
