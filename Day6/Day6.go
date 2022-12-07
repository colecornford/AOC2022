package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// wrong

func main() {
	part1()
	part2()
}

func getData() string {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(content)

}

func part1() {
	data := getData()
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	w1, w2 := 0, 14
	for w2 < len(data) {
		window := data[w1:w2]
		repeat := false

		for _, char := range alphabet {
			if strings.Count(window, string(char)) > 1 {
				repeat = true
				break
			}
		}
		if !repeat {
			fmt.Println(w2)
			break
		}
		w1++
		w2++
	}
}

func part2() {
	// data := getData()
}
