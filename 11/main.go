package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type monkey struct {
	id             int
	items          []int
	operationSign  string
	operationValue int
	divisible      int
	trueMonkey     int
	falseMonkey    int
	itemsInspected int
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	var monkeys []monkey

	for i := 0; i < len(lines); i += 7 {

		id, _ := strconv.Atoi(string(lines[i][7]))
		var items []int
		for _, it := range strings.Split(lines[i+1][18:], ", ") {
			item, _ := strconv.Atoi(it)
			items = append(items, item)
		}
		operationSign := string(lines[i+2][23])
		operationValue, _ := strconv.Atoi(lines[i+2][25:])
		divisible, _ := strconv.Atoi(lines[i+3][21:])
		trueMonkey, _ := strconv.Atoi(lines[i+4][29:])
		falseMonkey, _ := strconv.Atoi(lines[i+5][30:])
		monkeys = append(monkeys, monkey{
			id:             id,
			items:          items,
			operationSign:  operationSign,
			operationValue: operationValue,
			divisible:      divisible,
			trueMonkey:     trueMonkey,
			falseMonkey:    falseMonkey,
			itemsInspected: 0,
		})
	}

	monkeys1 := make([]monkey, len(monkeys))
	copy(monkeys1, monkeys)
	monkeys2 := make([]monkey, len(monkeys))
	copy(monkeys2, monkeys)

	// Part 1
	for round := 0; round < 20; round++ {
		for m := 0; m < len(monkeys1); m++ {
			mon := monkeys1[m]
			for _, item := range mon.items {
				operationValue := mon.operationValue
				if operationValue == 0 {
					operationValue = item
				}
				worryLevel := 0
				if mon.operationSign == "*" {
					worryLevel = item * operationValue
				}
				if mon.operationSign == "+" {
					worryLevel = item + operationValue
				}
				worryLevel = int(math.Floor(float64(worryLevel) / 3))
				if (worryLevel % mon.divisible) == 0 {
					monkeys1[mon.trueMonkey].items = append(monkeys1[mon.trueMonkey].items, worryLevel)
				} else {
					monkeys1[mon.falseMonkey].items = append(monkeys1[mon.falseMonkey].items, worryLevel)
				}
				monkeys1[m].itemsInspected = monkeys1[m].itemsInspected + 1
			}
			monkeys1[m].items = []int{}
		}
	}

	mostInspected := []int{0, 0}
	for _, mon := range monkeys1 {
		for k, most := range mostInspected {
			if mon.itemsInspected > most {
				mostInspected[k] = mon.itemsInspected
				break
			}
		}
	}
	fmt.Println(mostInspected[0] * mostInspected[1])

	/***********************************************/
	// Part 2
	var lcm = 1
	for _, m := range monkeys2 {
		lcm *= m.divisible
	}
	for round := 0; round < 10000; round++ {
		for m := 0; m < len(monkeys2); m++ {
			mon := monkeys2[m]
			for _, item := range mon.items {
				operationValue := mon.operationValue
				if operationValue == 0 {
					operationValue = item
				}
				worryLevel := 0
				if mon.operationSign == "*" {
					worryLevel = item * operationValue
				}
				if mon.operationSign == "+" {
					worryLevel = item + operationValue
				}
				worryLevel = worryLevel % lcm
				if (worryLevel % mon.divisible) == 0 {
					monkeys2[mon.trueMonkey].items = append(monkeys2[mon.trueMonkey].items, worryLevel)
				} else {
					monkeys2[mon.falseMonkey].items = append(monkeys2[mon.falseMonkey].items, worryLevel)
				}
				monkeys2[m].itemsInspected = monkeys2[m].itemsInspected + 1
			}
			monkeys2[m].items = []int{}
		}
	}
	mostInspected = []int{0, 0}
	for _, mon := range monkeys2 {
		for k, most := range mostInspected {
			if mon.itemsInspected > most {
				mostInspected[k] = mon.itemsInspected
				break
			}
		}
	}
	fmt.Println(mostInspected[0] * mostInspected[1])
}
