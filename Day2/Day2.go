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
	rounds := getData()

	var totalScore int = 0

	for _, round := range rounds {
		s := strings.Split(round, " ")
		totalScore += outcome(s[1], s[0])
	}

	fmt.Println(totalScore)
}

func part2() {
	rounds := getData()

	var totalScore int = 0

	for _, round := range rounds {
		s := strings.Split(round, " ")
		fmt.Print(s)
		totalScore += outcome2(s[1], s[0])
	}

	fmt.Println(totalScore)
}

func outcome(us string, them string) int {
	if us == "X" { // Rock
		switch them {
		case "A": // Rock
			return 4
		case "B": // Paper
			return 1
		case "C": // Scissors
			return 7
		}
	}

	if us == "Y" { // Paper
		switch them {
		case "A": // Rock
			return 8
		case "B": // Paper
			return 5
		case "C": // Scissors
			return 2
		}
	}

	if us == "Z" { // Scissors
		switch them {
		case "A": // Rock
			return 3
		case "B": // Paper
			return 9
		case "C": // Scissors
			return 6
		}
	}
	return 0
}

func outcome2(us string, them string) int {
	if us == "X" { // Loss
		switch them {
		case "A": // Rock
			return 3
		case "B": // Paper
			return 1
		case "C": // Scissors
			return 2
		}
	}

	if us == "Y" { // Draw
		switch them {
		case "A": // Rock
			return 4
		case "B": // Paper
			return 5
		case "C": // Scissors
			return 6
		}
	}

	if us == "Z" { // Win
		switch them {
		case "A": // Rock
			return 8
		case "B": // Paper
			return 9
		case "C": // Scissors
			return 7
		}
	}
	return 0
}
