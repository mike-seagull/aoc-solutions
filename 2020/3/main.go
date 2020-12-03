package main

import (
	"fmt"
	"os"
	"bufio"
)

func getInput(fileName string) []string {
	f, _ := os.Open(fileName)
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	result := []string{}
	// Use Scan.
	for scanner.Scan() {
		line := scanner.Text()
		// Append line to result.
		result = append(result, line)
	}
	return result
}
func logic(input []string, x int, y int) int64 {
	var hittrees int64 = 0
	j := 0
	i := 1
	if y % 2 == 0 {
		i++
	}
	for i<len(input) {
		j = j+x
		line := input[i]
		if j >= len(line) {
			j = j - len(line)
		}
		if string(line[j]) == "#" {
			hittrees++
		}
		i = i+y
	}
	return hittrees
}
func part1(input []string) {
	fmt.Printf("Part1: %d\n", logic(input, 3, 1))
}
func part2(input []string) {
	a := logic(input, 1, 1)
	b := logic(input, 3, 1)
	c := logic(input, 5, 1)
	d := logic(input, 7, 1)
	e := logic(input, 1, 2)
	fmt.Printf("Part2: %d\n", int64(a * b * c * d * e))
}
func main() {
//	fmt.Println("test")
//	part1(getInput("test_input.txt"))  // test input
//	part2(getInput("test_input.txt"))  // test input
	input := getInput("input.txt")
	part1(input)
	part2(input)
}
