package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Packet struct {
	Value int
	List  []Packet
}

func main() {
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")

	// Part1
	var pairs [][2]Packet

	for i := 0; i < len(lines); i += 2 {
		if lines[i] == "" {
			i--
			continue
		}

		var pair [2]Packet
		pair[0], _ = ParsePacket(lines[i], 0)
		pair[1], _ = ParsePacket(lines[i+1], 0)
		pairs = append(pairs, pair)
	}

	indices := 0
	for i, pair := range pairs {
		if Compare(pair[0], pair[1]) == 1 {
			indices += i + 1
		}
	}
	fmt.Println(indices)

	// Part2
	var packets []*Packet
	for _, line := range lines {
		if line == "" {
			continue
		}
		packet, _ := ParsePacket(line, 0)
		packets = append(packets, &packet)
	}

	divider1, _ := ParsePacket("[[2]]", 0)
	divider2, _ := ParsePacket("[[6]]", 0)
	packets = append(packets, &divider1, &divider2)

	sort.Slice(packets, func(i, j int) bool {
		return Compare(*packets[i], *packets[j]) == 1
	})

	key := 1
	for i, p := range packets {
		if p == &divider1 || p == &divider2 {
			key *= i + 1
		}
	}
	fmt.Println(key)
}

func ParsePacket(line string, start int) (Packet, int) {
	packet := Packet{
		Value: -1,
		List:  nil,
	}

	if line[start] != '[' {
		end := start
		for line[end] != ']' && line[end] != ',' {
			end++
		}

		packet.Value, _ = strconv.Atoi(line[start:end])
		return packet, end
	}

	start++
	for line[start] != ']' {
		if line[start] == ',' || line[start] == ' ' {
			start++
			continue
		}

		var child Packet
		child, start = ParsePacket(line, start)
		packet.List = append(packet.List, child)
	}

	return packet, start + 1
}

func Compare(left, right Packet) int {
	if left.Value >= 0 && right.Value >= 0 {
		if left.Value > right.Value {
			return -1
		}
		if left.Value < right.Value {
			return 1
		}
		return 0
	}

	if left.Value < 0 && right.Value < 0 {
		for i := 0; i < len(right.List) && i < len(left.List); i++ {
			v := Compare(left.List[i], right.List[i])
			if v != 0 {
				return v
			}
		}

		if len(left.List) < len(right.List) {
			return 1
		}
		if len(left.List) > len(right.List) {
			return -1
		}
		return 0
	}

	if left.Value >= 0 {
		packet := Packet{
			Value: -1,
			List:  []Packet{left},
		}
		return Compare(packet, right)
	}

	if right.Value >= 0 {
		packet := Packet{
			Value: -1,
			List:  []Packet{right},
		}
		return Compare(left, packet)
	}

	return 0
}
