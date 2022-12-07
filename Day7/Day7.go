package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// wrong 1046278 too low

func main() {
	part1()
	part2()
}

func getData() (lines []string) {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n")
}

func part1() {

	var fileSystem = make(map[string]int)
	var path string

	data := getData()
	for x, line := range data {
		if strings.Contains(line, "$ cd /") {
			path = ""
		} else if strings.Contains(line, "$ cd ..") {
			path = path[0:strings.LastIndex(path, "/")]
		} else if strings.Contains(line, "$ cd") {
			path += "/" + strings.Split(line, " ")[2]
		}
		if strings.Contains(line, "$ ls") {
			z := x + 1
			for {
				if z >= len(data) {
					break
				}
				if strings.Contains(data[z], "$") {
					break
				}
				if strings.Split(data[z], " ")[0] != "dir" {

					tmp, _ := strconv.Atoi(strings.Split(data[z], " ")[0])
					tmpPath := path
					for len(tmpPath) > 1 {
						fileSystem[tmpPath] += tmp
						tmpPath = tmpPath[0:strings.LastIndex(tmpPath, "/")]
					}
				}
				z++
			}
			fmt.Println(path)
		}
	}
	//fmt.Println(fileSystem)

	val := 0
	for _, dir := range fileSystem {
		if dir <= 100000 {
			val += dir
		}
	}
	fmt.Println(fileSystem)
	fmt.Println(val)
}

func part2() {

	var fileSystem = make(map[string]int)
	var path string

	data := getData()
	for x, line := range data {
		if strings.Contains(line, "$ cd /") {
			path = "/"
		} else if strings.Contains(line, "$ cd ..") {
			path = path[0:strings.LastIndex(path, "/")]
		} else if strings.Contains(line, "$ cd") {
			path += "/" + strings.Split(line, " ")[2]
		}
		if strings.Contains(line, "$ ls") {
			z := x + 1
			for {
				if z >= len(data) {
					break
				}
				if strings.Contains(data[z], "$") {
					break
				}
				if strings.Split(data[z], " ")[0] != "dir" {

					tmp, _ := strconv.Atoi(strings.Split(data[z], " ")[0])
					tmpPath := path
					for len(tmpPath) > 0 {
						fileSystem[tmpPath] += tmp
						tmpPath = tmpPath[0:strings.LastIndex(tmpPath, "/")]
					}
				}
				z++
			}
			fmt.Println(path)
		}
	}
	//fmt.Println(fileSystem)

	unused := 70000000 - fileSystem["/"]
	target := 30000000 - unused

	fmt.Println(fileSystem)
	fmt.Println(unused)
	fmt.Println(target)
	best := 70000000
	for _, item := range fileSystem {
		if item > target && item < best {
			best = item
		}
	}
	fmt.Println(best)
}
