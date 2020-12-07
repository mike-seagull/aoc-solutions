package main

import (
	"os"
	"bufio"
	"strings"
)

func getInput(fileName string) Graph {
	f, _ := os.Open(fileName)
	// Create new Scanner.
	scanner := bufio.NewScanner(f)

	graph := Graph{}

	for scanner.Scan() {
		line := scanner.Text() // "light red bags contain 1 bright white bag, 2 muted yellow bags."
		line = strings.TrimRight(line, ".") // "light red bags contain 1 bright white bag, 2 muted yellow bags"
		all_bags := strings.Split(line, "contain") // ["light red bags ", " 1 bright white bag, 2 muted yellow bags"]
		all_bags[1] = strings.TrimSpace(all_bags[1]) // "1 bright white bag, 2 muted yellow bags"
		contains_bags := strings.Split(all_bags[1], ",") // ["1 bright white bag"," 2 muted yellow bags"]
		for _, contains_bag := range contains_bags {
			contains_bag = strings.TrimSpace(contains_bag)
			bag_info := strings.SplitN(contains_bags, " ", 2)
			
		}
	}
	return result
}

func main() {
}