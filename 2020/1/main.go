package main
import (
	"fmt"
	"strconv"
	"bufio"
	"os"
)
func getInput(fileName string) []int {
	f, _ := os.Open(fileName)
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	result := []int{}
	// Use Scan.
	for scanner.Scan() {
		line := scanner.Text()
		// Append line to result.
		num, _ := strconv.Atoi(line)
		result = append(result, num)
	}
	return result
}
func part1(input []int) {
	for i:=0; i < len(input); i++ {
		for j:=i+1; j < len(input); j++ {
			if input[i] + input[j] == 2020 {
				fmt.Printf("Part1: %d * %d = %d\n", input[i], input[j], input[i]*input[j])
			}
		}
	}
}
func part2(input []int) {
	for i:=0; i < len(input); i++ {
		for j:=i+1; j < len(input); j++ {
			for k:=j+1; k < len(input); k++ {
				if input[i] + input[j] + input[k] == 2020 {
					fmt.Printf("Part2: %d * %d * %d = %d\n", input[i], input[j], input[k], input[i]*input[j]*input[k])
				}
			}
		}
	}
}
func main() {
	//testInput := []int {1721, 979, 366, 299, 675, 1456}
	//part1(testInput)
	input := getInput("input.txt")
	part1(input)
	part2(input)
}
