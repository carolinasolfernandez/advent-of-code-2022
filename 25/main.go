package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	sum := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		for i := 0; i < len(line); i++ {
			var num int
			switch line[len(line)-1-i] {
			case '2':
				num = 2
			case '1':
				num = 1
			case '0':
				num = 0
			case '-':
				num = -1
			case '=':
				num = -2
			}
			sum += int(math.Pow(5, float64(i))) * num
		}
	}

	part1 := ""
	for sum > 0 {
		switch sum % 5 {
		case 2:
			part1 = "2" + part1
		case 1:
			part1 = "1" + part1
		case 0:
			part1 = "0" + part1
		case 4:
			part1 = "-" + part1
			sum += 2
		case 3:
			part1 = "=" + part1
			sum += 3
		}
		sum /= 5
	}

	fmt.Println("Part 1:", part1)
}
