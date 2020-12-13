package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)
func getInput(fileName string) []int {
	f, _ := os.Open(fileName)
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	port_output := []int{}
	for scanner.Scan() {
		line := scanner.Text() // "nop +0"
		arg,_ := strconv.Atoi(line)
		port_output = append(port_output, arg)
	}
	return port_output
}
func part1(output []int, pre_len int) int{
	answer := -99
	for i:=pre_len; i<len(output);i++ {
		found := true
		for j:= i-1; j>= i-pre_len; j-- {
			for k:= i-1; k>= i-pre_len; k-- {
				if j != k && output[j] + output[k] == output[i] {
					found = false
				}
			}
		}
		if found {
			answer = output[i]
			break
		}
	}
	fmt.Printf("Part1: %d\n", answer)
	return answer
}
func part2(output []int, invalid int) {
	answer := -99
	for i:=0; i<len(output);i++ {
		n := output[i]
		max := n
		min := n
		found := false
		for j:=i+1; j<len(output);j++ {
			n += output[j]
			if output[j] < min {
				min = output[j]
			}
			if output[j] > max {
				max = output[j]
			}
			if n > invalid {
				break
			} else if n == invalid {
				found = true
				break
			}
		}
		if found {
			answer = max + min
			break
		}
	}
	fmt.Printf("Part2: %d", answer)
}
func main() {
	//ans := part1(getInput("test_input.txt"), 5)
	//part2(getInput("test_input.txt"), ans)
	input := getInput("input.txt")
	ans := part1(input, 25)
	part2(input, ans)
}
