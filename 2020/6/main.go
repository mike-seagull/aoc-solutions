package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
)
type Group map[string]int
func getInput(fileName string) []Group {
	f, _ := os.Open(fileName)
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	var result []Group
	// Use Scan.
	group := Group{} // { "question_letter": occurence_count }
	number_of_group_members := 0
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if len(line) == 0 { // end of passport
			group["number_of_group_members"] = number_of_group_members
			result = append(result, group)
			group = Group{}
			number_of_group_members = 0
		} else {
			for _, c := range line {
				group[string(c)] += 1
			}
			number_of_group_members++
		}
	}
	group["number_of_group_members"] = number_of_group_members
	result = append(result, group)
	return result
}
func part1(groups []Group) {
	sum_of_yes := 0
	for _, group := range groups {
		delete(group, "number_of_group_members")
		sum_of_yes += len(group)
	}
	fmt.Printf("Part1: %d\n", sum_of_yes)
}
func part2(groups []Group) {
	sum_of_yes := 0
	for _, group := range groups {
		number_of_group_members := group["number_of_group_members"]
		delete(group, "number_of_group_members")
		for _, val := range group {
			if val == number_of_group_members {
				sum_of_yes++
			}
		}
	}
	fmt.Printf("Part2: %d\n", sum_of_yes)
}
func main() {
	//part1(getInput("test_input.txt")) // test input
	//part2(getInput("test_input.txt")) // test input
	part1(getInput("input.txt"))
	part2(getInput("input.txt"))
}
