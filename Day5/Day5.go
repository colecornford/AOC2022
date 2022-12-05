package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

// wrong

func main() {
	part1()
	part2()
}

func getData() ([]string, []string) {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	tmp := strings.Split(string(content), "\n\n")
	return strings.Split(string(tmp[0]), "\n"), strings.Split(string(tmp[1]), "\n")

	// var stacks [2]string
}

func part1() {

	stackData, moveData := getData()
	fmt.Println(stackData)
	fmt.Println()
	fmt.Println(moveData)
	fmt.Println()

	var stack1, stack2, stack3, stack4, stack5, stack6, stack7, stack8, stack9, moves []string
	for _, line := range stackData {
		if string(line[1]) != (" ") && !strings.ContainsAny(string(line[1]), "0123456789") {
			stack1 = append(stack1, string(line[1]))
		}
		if string(line[5]) != (" ") && !strings.ContainsAny(string(line[5]), "0123456789") {
			stack2 = append(stack2, string(line[5]))
		}
		if string(line[9]) != (" ") && !strings.ContainsAny(string(line[9]), "0123456789") {
			stack3 = append(stack3, string(line[9]))
		}
		if string(line[13]) != (" ") && !strings.ContainsAny(string(line[13]), "0123456789") {
			stack4 = append(stack4, string(line[13]))
		}
		if string(line[17]) != (" ") && !strings.ContainsAny(string(line[17]), "0123456789") {
			stack5 = append(stack5, string(line[17]))
		}
		if string(line[21]) != (" ") && !strings.ContainsAny(string(line[21]), "0123456789") {
			stack6 = append(stack6, string(line[21]))
		}
		if string(line[25]) != (" ") && !strings.ContainsAny(string(line[25]), "0123456789") {
			stack7 = append(stack7, string(line[25]))
		}
		if string(line[29]) != (" ") && !strings.ContainsAny(string(line[29]), "0123456789") {
			stack8 = append(stack8, string(line[29]))
		}
		if string(line[33]) != (" ") && !strings.ContainsAny(string(line[33]), "0123456789") {
			stack9 = append(stack9, string(line[33]))
		}
	}
	fmt.Println(stack1)
	fmt.Println(stack2)
	fmt.Println(stack3)
	fmt.Println(stack4)
	fmt.Println(stack5)
	fmt.Println(stack6)
	fmt.Println(stack7)
	fmt.Println(stack8)
	fmt.Println(stack9)

	f := func(c rune) bool {
		return !unicode.IsNumber(c)
	}

	for _, line := range moveData {
		moves = strings.FieldsFunc(line, f)
		var x int
		x, _ = strconv.Atoi(moves[0])
		var toStack, fromStack []string
		if moves[1] == "1" {
			fromStack = stack1
		}
		if moves[1] == "2" {
			fromStack = stack2
		}
		if moves[1] == "3" {
			fromStack = stack3
		}
		if moves[1] == "4" {
			fromStack = stack4
		}
		if moves[1] == "5" {
			fromStack = stack5
		}
		if moves[1] == "6" {
			fromStack = stack6
		}
		if moves[1] == "7" {
			fromStack = stack7
		}
		if moves[1] == "8" {
			fromStack = stack8
		}
		if moves[1] == "9" {
			fromStack = stack9
		}

		if moves[2] == "1" {
			toStack = stack1
		}
		if moves[2] == "2" {
			toStack = stack2
		}
		if moves[2] == "3" {
			toStack = stack3
		}
		if moves[2] == "4" {
			toStack = stack4
		}
		if moves[2] == "5" {
			toStack = stack5
		}
		if moves[2] == "6" {
			toStack = stack6
		}
		if moves[2] == "7" {
			toStack = stack7
		}
		if moves[2] == "8" {
			toStack = stack8
		}
		if moves[2] == "9" {
			toStack = stack9
		}

		for x > 0 {
			toStack = append([]string{fromStack[0]}, toStack...)
			if len(fromStack) >= 1 {
				fromStack = fromStack[1:] // Pop
			}
			x--
		}

		if moves[1] == "1" {
			stack1 = fromStack
		}
		if moves[1] == "2" {
			stack2 = fromStack
		}
		if moves[1] == "3" {
			stack3 = fromStack
		}
		if moves[1] == "4" {
			stack4 = fromStack
		}
		if moves[1] == "5" {
			stack5 = fromStack
		}
		if moves[1] == "6" {
			stack6 = fromStack
		}
		if moves[1] == "7" {
			stack7 = fromStack
		}
		if moves[1] == "8" {
			stack8 = fromStack
		}
		if moves[1] == "9" {
			stack9 = fromStack
		}

		if moves[2] == "1" {
			stack1 = toStack
		}
		if moves[2] == "2" {
			stack2 = toStack
		}
		if moves[2] == "3" {
			stack3 = toStack
		}
		if moves[2] == "4" {
			stack4 = toStack
		}
		if moves[2] == "5" {
			stack5 = toStack
		}
		if moves[2] == "6" {
			stack6 = toStack
		}
		if moves[2] == "7" {
			stack7 = toStack
		}
		if moves[2] == "8" {
			stack8 = toStack
		}
		if moves[2] == "9" {
			stack9 = toStack
		}

		fmt.Println()
		fmt.Println(stack1)
		fmt.Println(stack2)
		fmt.Println(stack3)
		fmt.Println(stack4)
		fmt.Println(stack5)
		fmt.Println(stack6)
		fmt.Println(stack7)
		fmt.Println(stack8)
		fmt.Println(stack9)
	}

}

func part1Test() {

	stackData, moveData := getData()
	fmt.Println(stackData)
	fmt.Println()
	fmt.Println(moveData)
	fmt.Println()

	var stack1, stack2, stack3, moves []string
	for _, line := range stackData {
		if string(line[1]) != (" ") && !strings.ContainsAny(string(line[1]), "0123456789") {
			stack1 = append(stack1, string(line[1]))
		}
		if string(line[5]) != (" ") && !strings.ContainsAny(string(line[5]), "0123456789") {
			stack2 = append(stack2, string(line[5]))
		}
		if string(line[9]) != (" ") && !strings.ContainsAny(string(line[9]), "0123456789") {
			stack3 = append(stack3, string(line[9]))
		}
	}
	fmt.Println(stack1)
	fmt.Println(stack2)
	fmt.Println(stack3)

	f := func(c rune) bool {
		return !unicode.IsNumber(c)
	}

	for _, line := range moveData {
		moves = strings.FieldsFunc(line, f)
		var x int
		x, _ = strconv.Atoi(moves[0])
		var toStack, fromStack []string
		if moves[1] == "1" {
			fromStack = stack1
		}
		if moves[1] == "2" {
			fromStack = stack2
		}
		if moves[1] == "3" {
			fromStack = stack3
		}
		if moves[2] == "1" {
			toStack = stack1
		}
		if moves[2] == "2" {
			toStack = stack2
		}
		if moves[2] == "3" {
			toStack = stack3
		}
		for x > 0 {
			toStack = append([]string{fromStack[0]}, toStack...)
			if len(fromStack) >= 1 {
				fromStack = fromStack[1:] // Pop
			}
			x--
		}

		if moves[1] == "1" {
			stack1 = fromStack
		}
		if moves[1] == "2" {
			stack2 = fromStack
		}
		if moves[1] == "3" {
			stack3 = fromStack
		}
		if moves[2] == "1" {
			stack1 = toStack
		}
		if moves[2] == "2" {
			stack2 = toStack
		}
		if moves[2] == "3" {
			stack3 = toStack
		}

		fmt.Println()
		fmt.Println(stack1)
		fmt.Println(stack2)
		fmt.Println(stack3)
	}

}

func part2() {
	stackData, moveData := getData()
	fmt.Println(stackData)
	fmt.Println()
	fmt.Println(moveData)
	fmt.Println()

	var stack1, stack2, stack3, stack4, stack5, stack6, stack7, stack8, stack9, moves []string
	for _, line := range stackData {
		if string(line[1]) != (" ") && !strings.ContainsAny(string(line[1]), "0123456789") {
			stack1 = append(stack1, string(line[1]))
		}
		if string(line[5]) != (" ") && !strings.ContainsAny(string(line[5]), "0123456789") {
			stack2 = append(stack2, string(line[5]))
		}
		if string(line[9]) != (" ") && !strings.ContainsAny(string(line[9]), "0123456789") {
			stack3 = append(stack3, string(line[9]))
		}
		if string(line[13]) != (" ") && !strings.ContainsAny(string(line[13]), "0123456789") {
			stack4 = append(stack4, string(line[13]))
		}
		if string(line[17]) != (" ") && !strings.ContainsAny(string(line[17]), "0123456789") {
			stack5 = append(stack5, string(line[17]))
		}
		if string(line[21]) != (" ") && !strings.ContainsAny(string(line[21]), "0123456789") {
			stack6 = append(stack6, string(line[21]))
		}
		if string(line[25]) != (" ") && !strings.ContainsAny(string(line[25]), "0123456789") {
			stack7 = append(stack7, string(line[25]))
		}
		if string(line[29]) != (" ") && !strings.ContainsAny(string(line[29]), "0123456789") {
			stack8 = append(stack8, string(line[29]))
		}
		if string(line[33]) != (" ") && !strings.ContainsAny(string(line[33]), "0123456789") {
			stack9 = append(stack9, string(line[33]))
		}
	}
	fmt.Println(stack1)
	fmt.Println(stack2)
	fmt.Println(stack3)
	fmt.Println(stack4)
	fmt.Println(stack5)
	fmt.Println(stack6)
	fmt.Println(stack7)
	fmt.Println(stack8)
	fmt.Println(stack9)

	f := func(c rune) bool {
		return !unicode.IsNumber(c)
	}

	for _, line := range moveData {
		moves = strings.FieldsFunc(line, f)
		var x int
		x, _ = strconv.Atoi(moves[0])
		var toStack, fromStack []string
		if moves[1] == "1" {
			fromStack = stack1
		}
		if moves[1] == "2" {
			fromStack = stack2
		}
		if moves[1] == "3" {
			fromStack = stack3
		}
		if moves[1] == "4" {
			fromStack = stack4
		}
		if moves[1] == "5" {
			fromStack = stack5
		}
		if moves[1] == "6" {
			fromStack = stack6
		}
		if moves[1] == "7" {
			fromStack = stack7
		}
		if moves[1] == "8" {
			fromStack = stack8
		}
		if moves[1] == "9" {
			fromStack = stack9
		}

		if moves[2] == "1" {
			toStack = stack1
		}
		if moves[2] == "2" {
			toStack = stack2
		}
		if moves[2] == "3" {
			toStack = stack3
		}
		if moves[2] == "4" {
			toStack = stack4
		}
		if moves[2] == "5" {
			toStack = stack5
		}
		if moves[2] == "6" {
			toStack = stack6
		}
		if moves[2] == "7" {
			toStack = stack7
		}
		if moves[2] == "8" {
			toStack = stack8
		}
		if moves[2] == "9" {
			toStack = stack9
		}

		var newStack []string
		newStack = append(newStack, fromStack[0:x]...)
		for x > 0 {
			if len(fromStack) >= 0 {
				fromStack = fromStack[1:] // Pop
			}
			x--
		}

		toStack = append(newStack, toStack...)
		if moves[1] == "1" {
			stack1 = fromStack
		}
		if moves[1] == "2" {
			stack2 = fromStack
		}
		if moves[1] == "3" {
			stack3 = fromStack
		}
		if moves[1] == "4" {
			stack4 = fromStack
		}
		if moves[1] == "5" {
			stack5 = fromStack
		}
		if moves[1] == "6" {
			stack6 = fromStack
		}
		if moves[1] == "7" {
			stack7 = fromStack
		}
		if moves[1] == "8" {
			stack8 = fromStack
		}
		if moves[1] == "9" {
			stack9 = fromStack
		}

		if moves[2] == "1" {
			stack1 = toStack
		}
		if moves[2] == "2" {
			stack2 = toStack
		}
		if moves[2] == "3" {
			stack3 = toStack
		}
		if moves[2] == "4" {
			stack4 = toStack
		}
		if moves[2] == "5" {
			stack5 = toStack
		}
		if moves[2] == "6" {
			stack6 = toStack
		}
		if moves[2] == "7" {
			stack7 = toStack
		}
		if moves[2] == "8" {
			stack8 = toStack
		}
		if moves[2] == "9" {
			stack9 = toStack
		}

		fmt.Println()
		fmt.Println(stack1)
		fmt.Println(stack2)
		fmt.Println(stack3)
		fmt.Println(stack4)
		fmt.Println(stack5)
		fmt.Println(stack6)
		fmt.Println(stack7)
		fmt.Println(stack8)
		fmt.Println(stack9)
	}
}

func part2Test() {

	stackData, moveData := getData()
	fmt.Println(stackData)
	fmt.Println()
	fmt.Println(moveData)
	fmt.Println()

	var stack1, stack2, stack3, moves []string
	for _, line := range stackData {
		if string(line[1]) != (" ") && !strings.ContainsAny(string(line[1]), "0123456789") {
			stack1 = append(stack1, string(line[1]))
		}
		if string(line[5]) != (" ") && !strings.ContainsAny(string(line[5]), "0123456789") {
			stack2 = append(stack2, string(line[5]))
		}
		if string(line[9]) != (" ") && !strings.ContainsAny(string(line[9]), "0123456789") {
			stack3 = append(stack3, string(line[9]))
		}
	}
	fmt.Println(stack1)
	fmt.Println(stack2)
	fmt.Println(stack3)

	f := func(c rune) bool {
		return !unicode.IsNumber(c)
	}

	for _, line := range moveData {
		moves = strings.FieldsFunc(line, f)
		var x int
		x, _ = strconv.Atoi(moves[0])
		var toStack, fromStack []string
		if moves[1] == "1" {
			fromStack = stack1
		}
		if moves[1] == "2" {
			fromStack = stack2
		}
		if moves[1] == "3" {
			fromStack = stack3
		}
		if moves[2] == "1" {
			toStack = stack1
		}
		if moves[2] == "2" {
			toStack = stack2
		}
		if moves[2] == "3" {
			toStack = stack3
		}
		var newStack []string
		newStack = append(newStack, fromStack[0:x]...)
		for x > 0 {
			if len(fromStack) >= 0 {
				fromStack = fromStack[1:] // Pop
			}
			x--
		}

		toStack = append(newStack, toStack...)
		if moves[1] == "1" {
			stack1 = fromStack
		}
		if moves[1] == "2" {
			stack2 = fromStack
		}
		if moves[1] == "3" {
			stack3 = fromStack
		}
		if moves[2] == "1" {
			stack1 = toStack
		}
		if moves[2] == "2" {
			stack2 = toStack
		}
		if moves[2] == "3" {
			stack3 = toStack
		}

		fmt.Println()
		fmt.Println(stack1)
		fmt.Println(stack2)
		fmt.Println(stack3)
	}

}
