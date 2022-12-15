package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type sensorType struct {
	x      float64
	y      float64
	beacon struct {
		x float64
		y float64
	}
}

func main() {
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	var sensors []sensorType
	maxX := float64(0)
	minX := float64(math.MaxInt)

	// set sensors and beacons
	for _, line := range lines {
		if line == "" {
			break
		}

		sensor := sensorType{}
		var pos int
		sensor.x, pos = getValue(0, line)
		sensor.y, pos = getValue(pos, line)
		if sensor.x > maxX {
			maxX = sensor.x
		}
		if sensor.x < minX {
			minX = sensor.x
		}
		sensor.beacon.x, pos = getValue(pos, line)
		sensor.beacon.y, pos = getValue(pos, line)
		if sensor.beacon.x > maxX {
			maxX = sensor.beacon.x
		}
		if sensor.beacon.x < minX {
			minX = sensor.beacon.x
		}

		if pos == 0 {
			panic("format error")
		}
		sensors = append(sensors, sensor)
	}

	// Part 1
	for _, sensor := range sensors {
		dist := math.Abs(sensor.x-sensor.beacon.x) + math.Abs(sensor.y-sensor.beacon.y)

		if sensor.x-dist < minX {
			minX = sensor.x - dist
		}
		if sensor.x+dist > maxX {
			maxX = sensor.x + dist
		}
	}

	positions := 0
	for x := minX; x <= maxX; x++ {
		covered, _ := positionCoveredBySensor(sensors, x, 2000000)
		if covered {
			positions++
		}
	}
	fmt.Println(positions)

	// Part 2
	found := false
	var tuning float64
	for y := float64(0); y <= 4000000; y++ {
		for x := float64(0); x <= 4000000; x++ {
			covered, newX := positionCoveredBySensor(sensors, x, y)
			if covered == false {
				found = true
				tuning = x*4000000 + y
				break
			}
			x = newX
		}
		if found {
			break
		}
	}
	if found == false {
		panic("No tuning found")
	}
	fmt.Println(int(tuning))

}

func positionCoveredBySensor(sensors []sensorType, x, y float64) (bool, float64) {
	for _, sensor := range sensors {
		if sensor.beacon.x == x && sensor.beacon.y == y {
			return false, x
		}
		dist := math.Abs(sensor.x-sensor.beacon.x) + math.Abs(sensor.y-sensor.beacon.y)
		if sensor.y-dist > y || sensor.y+dist < y {
			continue
		}

		overlap := dist - math.Abs(y-sensor.y)
		if (sensor.x-overlap) <= x && (sensor.x+overlap >= x) {
			return true, sensor.x + overlap
		}
	}
	return false, x
}

func getValue(init int, line string) (float64, int) {
	for i := init; i < len(line); i++ {
		if line[i] != '=' {
			continue
		}
		for j := i + 1; j < len(line); j++ {
			if j == len(line)-1 || line[j] == ',' || line[j] == ':' {
				var value int
				var err error
				if j == len(line)-1 {
					value, err = strconv.Atoi(line[i+1:])
				} else {
					value, err = strconv.Atoi(line[i+1 : j])
				}
				if err != nil {
					panic("conversion error")
				}
				return float64(value), j
			}
		}
	}
	return 0, 0
}
