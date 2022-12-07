package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	marker1 := 4
	marker2 := 14
	part1 := getAmountProcessed(content, marker1)
	part2 := getAmountProcessed(content, marker2)

	fmt.Println(part1 + marker1)
	fmt.Println(part2 + marker2)
}

func getAmountProcessed(content []byte, marker int) int {
	i := 0
	for {
		letters := content[i : i+marker]
		cont := false
		let := map[string]int{}

		for j := 0; j < marker; j++ {
			l1 := string(letters[j])
			val := let[l1]
			if val == 1 {
				cont = true
				break
			} else {
				let[l1] = 1
			}
		}

		if cont == false {
			break
		}
		i++
	}
	return i
}
