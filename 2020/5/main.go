package main

import (
	"os"
	"bufio"
	"fmt"
)
func getInput(fileName string) []string {
	f, _ := os.Open(fileName)
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	result := []string{}
	// Use Scan.
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return result
}
func find(s string, i int, low int, high int) int {
	diff := ((high + 1) - low) / 2
	if string(s[i]) == "F" || string(s[i]) == "L" {
		if i < len(s) - 1 {
			high -= diff
			find(s, i+1, low, high)
		} else {
			return low
		}
	} else if string(s[i]) == "B" || string(s[i]) == "R" {
		if i < len(s) - 1 {
			low += diff
			find(s, i+1, low, high)
		} else {
			return high
		}
	}
	return find(s, i+1, low, high)
}
func part1(seats []string) int{
	max_id := 0
	for _, seat := range seats {
		row := find(seat[:7], 0, 0, 127)
		col := find(seat[7:len(seat)], 0, 0, 7)
		seat_id := (row * 8) + col
		if seat_id > max_id {
			max_id = seat_id
		}
	}
	fmt.Printf("Part1: %d\n", max_id)
	return max_id
}
func part2(seats []string, max_id int) {
	seat_assignments := make([]int, max_id+1)
	for _, seat := range seats {
		row := find(seat[:7], 0, 0, 127)
		col := find(seat[7:len(seat)], 0, 0, 7)
		seat_id := (row * 8) + col
		seat_assignments[seat_id] = 1
	}
	for i, seat := range seat_assignments {
		if i != 0 && i < len(seat_assignments)-1 && seat == 0 && seat_assignments[i-1] == 1 && seat_assignments[i+1] == 1 {
			fmt.Printf("Part2: %d\n", i)
			break
		}
	}
}
func main() {
	//part1(getInput("test_input.txt")) // test input
	max_id := part1(getInput("input.txt"))
	//part2(getInput("test_input2.txt")) // test input 2
	part2(getInput("input.txt"), max_id)
}
