package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// wrong 604, 586, 589

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

	sections := getData()
	var total int = 0
	for _, section := range sections {

		group1 := make(map[int]int)
		group2 := make(map[int]int)

		var pair1, pair2 string = strings.Split(section, ",")[0], strings.Split(section, ",")[1]
		start1, _ := strconv.Atoi(strings.Split(pair1, "-")[0])
		end1, _ := strconv.Atoi(strings.Split(pair1, "-")[1])
		for start1 <= end1 {
			group1[start1] = start1
			start1++
		}
		start2, _ := strconv.Atoi(strings.Split(pair2, "-")[0])
		end2, _ := strconv.Atoi(strings.Split(pair2, "-")[1])
		for start2 <= end2 {
			group2[start2] = start2
			start2++
		}

		fmt.Println(group1)
		fmt.Println(group2)

		var sub1, sub2 bool = true, true

		for _, item := range group1 {
			if group2[item] != item {
				sub1 = false
			}
		}

		for _, item := range group2 {
			if group1[item] != item {
				sub2 = false
			}
		}

		if sub1 || sub2 {
			total++
			fmt.Println(total)
		}
	}
	fmt.Println(total)
}

func part2() {

	sections := getData()
	var total int = 0
	for _, section := range sections {

		group1 := make(map[int]int)
		group2 := make(map[int]int)

		var pair1, pair2 string = strings.Split(section, ",")[0], strings.Split(section, ",")[1]
		start1, _ := strconv.Atoi(strings.Split(pair1, "-")[0])
		end1, _ := strconv.Atoi(strings.Split(pair1, "-")[1])
		for start1 <= end1 {
			group1[start1] = start1
			start1++
		}
		start2, _ := strconv.Atoi(strings.Split(pair2, "-")[0])
		end2, _ := strconv.Atoi(strings.Split(pair2, "-")[1])
		for start2 <= end2 {
			group2[start2] = start2
			start2++
		}

		fmt.Println(group1)
		fmt.Println(group2)

		var sub1 bool = false

		for _, item := range group1 {
			if group2[item] == item {
				sub1 = true
			}
		}

		if sub1 {
			total++
			fmt.Println(total)
		}
	}
	fmt.Println(total)

}
