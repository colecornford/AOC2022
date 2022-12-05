package main

import (
	"fmt"
	"io/ioutil"
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
	rucksacks := getData()

	var priorities string = " abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // Position == Priority
	var sumPriorities int = 0
	for _, rucksack := range rucksacks {
		var compartment1 string = rucksack[0 : len(rucksack)/2]
		var compartment2 string = rucksack[len(rucksack)/2 : len(rucksack)]

		for _, item := range compartment1 {
			if strings.Contains(compartment2, string(item)) {
				sumPriorities += strings.Index(priorities, string(item))
				break
			}
		}
	}

	fmt.Println(sumPriorities)
}

func part2() {
	rucksacks := getData()

	var priorities string = " abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // Position == Priority
	var sumPriorities int = 0
	x := 0
	for x < len(rucksacks) { // Read 3 Lines
		var sack1, sack2, sack3 string = rucksacks[x], rucksacks[x+1], rucksacks[x+2]
		for _, item := range sack1 {
			if strings.Contains(sack2, string(item)) && strings.Contains(sack3, string(item)) {
				sumPriorities += strings.Index(priorities, string(item))
				x = x + 3
				break
			}
		}
	}
	fmt.Println(sumPriorities)
}
