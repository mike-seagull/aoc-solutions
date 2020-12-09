package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)
type Instruction struct {
	Operation string
	Argument int
}
func getInput(fileName string) []Instruction {
	f, _ := os.Open(fileName)
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	instructions := []Instruction{}
	for scanner.Scan() {
		line := scanner.Text() // "nop +0"
		line_split := strings.Split(line, " ") // ["nop", "+0"]
		arg,_ := strconv.Atoi(line_split[1])
		instructions = append(instructions, Instruction{line_split[0], arg})
	}
	return instructions
}
func check(instructions []Instruction) (bool, int){
	visited := make(map[int]int)
	accumulator := 0
	for i := 0; i != len(instructions); i++ {
		instruction := instructions[i]
		if _, found := visited[i]; found {
			return false, accumulator
			break
		} else {
			visited[i] = 1
			if instruction.Operation == "acc"{
				accumulator += instruction.Argument
			} else if instruction.Operation == "jmp" && instruction.Argument == 0 {
				return false, accumulator
			} else if instruction.Operation == "jmp" {
				i += (instruction.Argument - 1)
			}
		}
	}
	return true, accumulator
}
func part1(instructions []Instruction) {
	_, accumulator := check(instructions)
	fmt.Printf("Part1: %d", accumulator)
}
func part2(instructions []Instruction) {
	for i, instruction := range instructions {
		if instruction.Operation == "nop"{
			var new_instructs = make([]Instruction, len(instructions))
			copy(new_instructs, instructions)
			new_instruct := Instruction{"jmp", instruction.Argument}
			new_instructs[i] = new_instruct
			success, accumulator := check(new_instructs)
			if success {
				fmt.Printf("Part2: %d", accumulator)
				break
			}
		} else if instruction.Operation == "jmp"{
			var new_instructs = make([]Instruction, len(instructions))
			copy(new_instructs, instructions)
			new_instruct := Instruction{"nop", instruction.Argument}
			new_instructs[i] = new_instruct
			success, accumulator := check(new_instructs)
			if success {
				fmt.Printf("Part2: %d", accumulator)
				break
			}
		}
	}
}
func main() {
	//part1(getInput("test_input.txt"))
	//part2(getInput("test_input.txt"))
	input := getInput("input.txt")
	part1(input)
	part2(input)
}
