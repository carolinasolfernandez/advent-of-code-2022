package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type monkey struct {
	name      string
	number    *int
	monkeys   []string
	operation string
}

func main() {
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	monkeys := map[string]monkey{}
	for _, line := range lines {
		if line == "" {
			continue
		}

		m := monkey{}
		name := line[0:4]
		if _, ok := monkeys[name]; ok == true {
			panic("monkey already parsed")
		}
		m.name = name

		yell := line[6:]
		num, err := strconv.Atoi(yell)
		if err == nil {
			m.number = &num
		} else {
			data := strings.Split(yell, " ")
			if data[1] != "+" && data[1] != "-" && data[1] != "*" && data[1] != "/" {
				panic("format error")
			}
			m.operation = data[1]
			m.monkeys = []string{data[0], data[2]}
		}
		monkeys[name] = m
	}
	
	// Part 1
	root := monkeys["root"]
	var yell *int
	if root.number == nil {
		monkeys, yell = getNumber(monkeys, root)
	}
	fmt.Println("Part 1", *yell)

	// Part 2
	root = monkeys["root"]
	m1 := monkeys[root.monkeys[0]]
	m2 := monkeys[root.monkeys[1]]
	var v int
	var ok bool
	if v, ok = getNumberInverse(monkeys, m1, *m2.number); ok == false {
		v, _ = getNumberInverse(monkeys, m2, *m1.number)
	}
	fmt.Println("Part 2", v)

}

func getNumber(monkeys map[string]monkey, m monkey) (map[string]monkey, *int) {
	if m.number != nil {
		return monkeys, m.number
	}
	m1 := monkeys[m.monkeys[0]]
	m2 := monkeys[m.monkeys[1]]
	var n1, n2 *int
	if m1.number == nil {
		monkeys, m1.number = getNumber(monkeys, m1)
	}
	n1 = m1.number
	if m2.number == nil {
		monkeys, m2.number = getNumber(monkeys, m2)
	}
	n2 = m2.number
	var m3 int
	switch m.operation {
	case "+":
		m3 = *n1 + *n2
	case "-":
		m3 = *n1 - *n2
	case "*":
		m3 = *n1 * *n2
	case "/":
		m3 = *n1 / *n2
	}
	m.number = &m3
	monkeys[m.name] = m
	return monkeys, m.number
}

func getNumberInverse(monkeys map[string]monkey, m monkey, wanted int) (int, bool) {
	var newWanted int
	if m.name == "humn" {
		return wanted, true
	}

	if m.monkeys == nil {
		return 0, false
	}

	m1 := monkeys[m.monkeys[0]]
	m2 := monkeys[m.monkeys[1]]

	switch m.operation {
	case "+":
		newWanted = wanted - *m2.number
	case "-":
		newWanted = wanted + *m2.number
	case "*":
		newWanted = wanted / *m2.number
	case "/":
		newWanted = wanted * *m2.number
	}

	if v, ok := getNumberInverse(monkeys, m1, newWanted); ok {
		return v, true
	}

	switch m.operation {
	case "+":
		newWanted = wanted - *m1.number
	case "-":
		newWanted = -wanted + *m1.number
	case "*":
		newWanted = wanted / *m1.number
	case "/":
		newWanted = wanted * *m1.number
	}

	return getNumberInverse(monkeys, m2, newWanted)
}
