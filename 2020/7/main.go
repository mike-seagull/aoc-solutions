package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func getInput(fileName string) map[string][]string {
	f, _ := os.Open(fileName)
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	nodes := map[string][]string{}
	for scanner.Scan() {
		line := scanner.Text() // "light red bags contain 1 bright white bag, 2 muted yellow bags."
		line = strings.TrimRight(line, ".") // "light red bags contain 1 bright white bag, 2 muted yellow bags"
		all_bags := strings.Split(line, "contain") // ["light red bags ", " 1 bright white bag, 2 muted yellow bags"]
		all_bags[0] = strings.TrimSpace(all_bags[0]) // "light red bags"
		all_bags[1] = strings.TrimSpace(all_bags[1]) // "1 bright white bag, 2 muted yellow bags"
		contains_bags := strings.Split(all_bags[1], ",") // ["1 bright white bag"," 2 muted yellow bags"]
		for _, contains_bag := range contains_bags {
			contains_bag = strings.TrimSpace(contains_bag)
			bag_info := strings.SplitN(contains_bag, " ", 2) // ["1", "bright white bag"]
			if _, ok := nodes[bag_info[1]]; !ok {
				nodes[bag_info[1]] = []string{}
			}
			nodes[bag_info[1]] = append(nodes[bag_info[1]], all_bags[0])
		}
	}
	return nodes
}
func part1(nodes map[string][]string, bag string, visited map[string]int) int {
	count := 0
	for _, bag_name := range nodes[bag] {
		if _, ok := visited[bag_name]; !ok {
			visited[bag_name] = 1
			count++
			return part1(nodes, bag_name, visited)
		}
	}
	return count
}
func main() {
	c := part1(getInput("test_input.txt"), "shiny gold bag", make(map[string]int))
	fmt.Println(c)
}
