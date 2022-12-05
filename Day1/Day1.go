package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

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
	rations := getData()

	elves := make(map[int]int)
	var elf, totalRation int = 1, 0

	for _, ration := range rations {
		if ration == "" {
			elves[elf] = totalRation
			elf++
			totalRation = 0
		} else {
			x, _ := strconv.Atoi(ration)
			totalRation += x
		}
	}

	max := 0
	for _, item := range elves {
		if item > max {
			max = item
		}
	}
	fmt.Println(elves)
	fmt.Println(max)
}

func part2() {
	rations := getData()

	elves := make(map[int]int)
	var elf, totalRation int = 1, 0

	for _, ration := range rations {
		if ration == "" {
			elves[elf] = totalRation
			elf++
			totalRation = 0
		} else {
			x, _ := strconv.Atoi(ration)
			totalRation += x
		}
	}

	first, second, third := 0, 0, 0
	for _, item := range elves {
		if item > first {
			third = second
			second = first
			first = item
		} else if item > second {
			third = second
			second = item
		} else if item > third {
			third = item
		}
	}
	fmt.Println(elves)
	fmt.Println(first + second + third)
}
