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

	trees := getData()
	insidePoints, outsidePoints := make(map[string]bool), make(map[string]bool)
	count := 0
	for y, line := range trees {
		for x, h := range line {
			var visible bool = true
			var curTreeHeight, compTreeHeight int
			curTreeHeight, _ = strconv.Atoi(string(h))
			var tmp string = strconv.Itoa(x) + "," + strconv.Itoa(y)
			if _, ok := insidePoints[tmp]; ok {
				continue
			}
			if _, ok := outsidePoints[tmp]; ok {
				continue
			}
			if y == 0 || x == 0 || y == len(line)-1 || x == len(trees)-1 {
				outsidePoints[tmp] = true
				continue
				// This is for outside ring
			}
			for z := x + 1; z < len(line); z++ {
				compTreeHeight, _ = strconv.Atoi(string(line[z]))
				fmt.Printf("\n %d vs %d \n", compTreeHeight, curTreeHeight)
				if compTreeHeight >= curTreeHeight {
					visible = false
					break
				}
			}
			if visible {
				insidePoints[tmp] = true
				continue
			} else {
				visible = true
			}
			for z := x - 1; z >= 0; z-- {
				compTreeHeight, _ = strconv.Atoi(string(line[z]))
				if compTreeHeight >= curTreeHeight {
					visible = false
					break
				}
			}
			if visible {
				insidePoints[tmp] = true
				continue
			} else {
				visible = true
			}
			for z := y + 1; z < len(trees); z++ {
				compTreeHeight, _ = strconv.Atoi(string(trees[z][x]))
				if compTreeHeight >= curTreeHeight {
					visible = false
					break
				}
			}
			if visible {
				insidePoints[tmp] = true
				continue
			} else {
				visible = true
			}
			for z := y - 1; z >= 0; z-- {
				compTreeHeight, _ = strconv.Atoi(string(trees[z][x]))
				fmt.Print(compTreeHeight)
				if compTreeHeight >= curTreeHeight {
					visible = false
					break
				}
			}
			if visible {
				insidePoints[tmp] = true
				continue
			}
		}
	}
	for _, item := range outsidePoints {
		if item {
			count++
		}
	}
	for _, item := range insidePoints {
		if item {
			count++
		}
	}
	fmt.Println(outsidePoints)
	fmt.Println()
	fmt.Println(insidePoints)
	fmt.Println()
	fmt.Println(count)
}

func part2() {

	trees := getData()
	scores := make(map[string]int)
	for y, line := range trees {
		for x, h := range line {
			viewDistance := 0
			if y == 0 || x == 0 || y == len(line)-1 || x == len(trees)-1 {
				continue
				// Skip outside ring
			}
			var currentScore []int
			var curTreeHeight, compTreeHeight int
			curTreeHeight, _ = strconv.Atoi(string(h))
			var tmp string = strconv.Itoa(x) + "," + strconv.Itoa(y)
			for z := x + 1; z < len(line); z++ {
				compTreeHeight, _ = strconv.Atoi(string(line[z]))
				if curTreeHeight > compTreeHeight {
					viewDistance++
				} else {
					viewDistance++
					break
				}
			}
			currentScore = append(currentScore, viewDistance)
			viewDistance = 0

			for z := x - 1; z >= 0; z-- {
				compTreeHeight, _ = strconv.Atoi(string(line[z]))
				if curTreeHeight > compTreeHeight {
					viewDistance++
				} else {
					viewDistance++
					break
				}
			}
			currentScore = append(currentScore, viewDistance)
			viewDistance = 0

			for z := y + 1; z < len(trees); z++ {
				compTreeHeight, _ = strconv.Atoi(string(trees[z][x]))
				if curTreeHeight > compTreeHeight {
					viewDistance++
				} else {
					viewDistance++
					break
				}
			}
			currentScore = append(currentScore, viewDistance)
			viewDistance = 0

			for z := y - 1; z >= 0; z-- {
				compTreeHeight, _ = strconv.Atoi(string(trees[z][x]))
				if curTreeHeight > compTreeHeight {
					viewDistance++
				} else {
					viewDistance++
					break
				}
			}
			currentScore = append(currentScore, viewDistance)
			fmt.Println(currentScore)
			var scenicScore = currentScore[0] * currentScore[1] * currentScore[2] * currentScore[3]
			fmt.Println(scenicScore)
			scores[tmp] = scenicScore
		}
	}
	fmt.Println(scores)

	var highest int = 0
	for _, i := range scores {
		if i > highest {
			highest = i
		}
	}
	fmt.Println(highest)
}
