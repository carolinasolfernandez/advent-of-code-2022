package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type valve struct {
	name    string
	rate    int
	tunnels []string
}

func main() {
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")

	valves := make(map[string]valve, len(lines))
	shortest := map[string]int{}
	distance := 1

	// Set valves
	for _, line := range lines {
		if line == "" {
			break
		}
		values := regexp.MustCompile("Valve | has flow rate=|; tunnels lead to valves |\n").Split(line, -1)
		name := values[1]
		rate, err := strconv.Atoi(values[2])
		if err != nil {
			values = regexp.MustCompile("Valve | has flow rate=|; tunnel leads to valve |\n").Split(line, -1)
			name = values[1]
			rate, err = strconv.Atoi(values[2])
			if err != nil {
				panic("conversion error")
			}
		}
		tunnels := strings.Split(values[3], ", ")
		for _, tunnel := range tunnels {
			shortest[name+tunnel] = distance
			shortest[tunnel+name] = distance
		}

		newValve := valve{
			name:    name,
			rate:    rate,
			tunnels: tunnels,
		}
		valves[name] = newValve
	}

	// Set shortest distance
	cont := true
	for cont {
		cont = false
		distance++
		for k, _ := range valves {
			for k2, val2 := range valves {
				if _, ok := shortest[k+k2]; ok {
					continue
				}

				for _, tunnel := range val2.tunnels {

					if dist, ok := shortest[k+tunnel]; ok && dist < distance {
						cont = true
						shortest[k+k2] = distance
						shortest[k2+k] = distance
						break
					}
				}
			}
		}
	}

	// Save valves to visit
	toVisit := map[string]bool{}
	for _, val := range valves {
		if val.rate > 0 {
			toVisit[val.name] = false
		}
	}

	// Part 1
	println(calculatePart1(valve{name: "AA"}, 31, shortest, valves, toVisit))

	// Part 2
	println(calculatePart2(valve{name: "AA"}, valve{name: "AA"}, 26, 27, shortest, valves, toVisit))
}

func calculatePart1(v valve, time int, shortest map[string]int, valves map[string]valve, toVisit map[string]bool) int {
	time--
	if time <= 0 {
		return 0
	}
	max := 0
	for k, visit := range toVisit {
		if visit == true {
			continue
		}
		toVisit[k] = true
		release := calculatePart1(valves[k], time-shortest[v.name+k], shortest, valves, toVisit)
		if release > max {
			max = release
		}
		toVisit[k] = false
	}
	return max + v.rate*time
}

func calculatePart2(v, v2 valve, time, time2 int, shortest map[string]int, valves map[string]valve, toVisit map[string]bool) int {
	time2--
	if time2 <= 0 {
		return 0
	}
	max := 0
	for k, visit := range toVisit {
		if visit == true {
			continue
		}
		toVisit[k] = true
		release := calculatePart2(v2, valves[k], time2, time-shortest[v.name+k], shortest, valves, toVisit)
		if release > max {
			max = release
		}
		toVisit[k] = false
	}
	return max + v2.rate*time2
}
