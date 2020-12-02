package main

import (
	"fmt"
	"regexp"
	"os"
	"bufio"
	"strconv"
	"strings"
)

type Password struct {
	Min int
	Max int
	Char string
	Value string
}
func getInput(fileName string) []Password {
	f, _ := os.Open(fileName)
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	result := []Password{}
	// Use Scan.
	for scanner.Scan() {
		line := scanner.Text()
		// Append line to result.
		p := Password{}
		s := strings.Fields(line)  // ["1-3", "a:", "abcde"]
		n := strings.Split(s[0], "-")  // ["1", "3"]
		p.Min, _ = strconv.Atoi(n[0])
		p.Max, _ = strconv.Atoi(n[1])
		p.Char = string(s[1][0])
		p.Value = s[2]
		result = append(result, p)
	}
	return result
}
func part1(passwords []Password) {
	validPasswords := 0
	for _, password := range passwords {
		reg := regexp.MustCompile(password.Char)
		matches := reg.FindAllStringIndex(password.Value, -1)
		if len(matches) >= password.Min && len(matches) <= password.Max {
			validPasswords++
		}
	}
	fmt.Printf("Part1: %d\n", validPasswords)
}
func part2(passwords []Password) {
	validPasswords := 0
	for _, password := range passwords {
		if (string(password.Value[password.Min-1]) == password.Char || string(password.Value[password.Max-1]) == password.Char) && !(string(password.Value[password.Min-1]) == password.Char && string(password.Value[password.Max-1]) == password.Char) {
			validPasswords++
		}
	}
	fmt.Printf("Part2: %d\n", validPasswords)
}
func main() {
	//testInput := []Password {Password {1, 3, "a", "abcde"}, Password {1, 3, "b", "cdefg"}, Password {2, 9, "c", "ccccccccc"}}
	// part1(testInput)
	//part2(testInput)
	input := getInput("input.txt")
	part1(input)
	part2(input)
}
