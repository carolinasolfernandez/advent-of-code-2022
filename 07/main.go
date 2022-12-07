package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Folder struct {
	name     string
	size     int
	files    []string
	children []*Folder
	parent   *Folder
}

var sizes int
var minPossible int
var free int

func main() {
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	root := &Folder{
		"root",
		0,
		[]string{},
		nil,
		nil,
	}
	currentFolder := root
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			break
		}

		if string(line[0:4]) == "$ cd" {
			val := strings.Split(line, " ")
			if val[2] == "/" {
				currentFolder = root
			} else if val[2] == ".." {
				currentFolder = currentFolder.parent
			} else {
				found := false
				for _, v := range currentFolder.children {
					if v.name == val[2] {
						found = true
						currentFolder = v
					}
				}
				if found == false {
					panic("no folder")
				}
			}
		} else if string(line) == "$ ls" {

		} else if string(line[0:3]) == "dir" {
			val := strings.Split(line, " ")
			found := false
			for _, v := range currentFolder.children {
				if v.name == val[1] {
					found = true
				}
			}
			if found == false {
				folder := &Folder{
					val[1],
					0,
					[]string{},
					nil,
					currentFolder,
				}
				currentFolder.children = append(currentFolder.children, folder)
			}

		} else {
			val := strings.Split(line, " ")
			size, err := strconv.Atoi(val[0])
			if err != nil {
				panic("conversion error")
			}
			currentFolder.size += size
			currentFolder.files = append(currentFolder.files, val[1])
			foldIt := currentFolder
			for foldIt.parent != nil {
				foldIt.parent.size += size
				foldIt = foldIt.parent
			}

		}
	}

	// Part 1
	sizes = 0
	if root.size < 100000 {
		sizes = root.size
	}
	getSize(root)
	fmt.Println(sizes)

	// Part 2
	total := 70000000
	needed := 30000000
	free = root.size - (total - needed)

	minPossible = root.size
	findMinPossible(root)

	fmt.Println(minPossible)
}

func getSize(root *Folder) *Folder {
	if root == nil {
		return nil
	}
	for _, v := range root.children {
		if v.size < 100000 {
			sizes = sizes + v.size
		}
		getSize(v)
	}
	return nil
}

func findMinPossible(root *Folder) *Folder {
	if root == nil {
		return nil
	}
	for _, v := range root.children {
		if v.size >= free && v.size <= minPossible {
			minPossible = v.size
		}
		findMinPossible(v)
	}
	return nil
}
