package main

import (
	"fmt"
	"os"
	"bufio"
)
type fn func(int, int, []string) int

func getInput(fileName string) []string {
	f, _ := os.Open(fileName)
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	seating := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		seating = append(seating, line)
	}
	return seating
}
func get_adjacent_occupied_seats(i int, j int, seating []string) int {
	occupied_seats := 0
	if i > 0 {
		if string(seating[i-1][j]) == "#" {
			occupied_seats++
		}
		if j > 0 {
			if string(seating[i-1][j-1]) == "#" {
				occupied_seats++
			}
		}
		if j < len(string(seating[i])) - 1 {
			if string(seating[i-1][j+1]) == "#" {
				occupied_seats++
			}
		}
	}
	if i < len(seating) - 1 {
		if string(seating[i+1][j]) == "#" {
			occupied_seats++
		}
		if j > 0 {
			if string(seating[i+1][j-1]) == "#" {
				occupied_seats++
			}
		}
		if j < len(string(seating[i])) - 1 {
			if string(seating[i+1][j+1]) == "#" {
				occupied_seats++
			}
		}
	}
	if j > 0 {
		if string(seating[i][j-1]) == "#" {
			occupied_seats++
		}
	}
	if j < len(string(seating[i])) - 1 {
		if string(seating[i][j+1]) == "#" {
			occupied_seats++
		}
	}
	return occupied_seats
}
func get_visibly_occupied_seats(i int, j int, seating []string) int {
	occupied_seats := 0
	derivatives := [][]int {{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for _, derivative := range derivatives {
		nx := i + derivative[0]
		ny := j + derivative[1]
		for ((0 <= nx) && (nx <= len(seating)-1)) && ((0 <= ny) && (ny <= len(string(seating[i]))-1)) {
			if string(seating[nx][ny]) == "L" {
				break
			}
			if string(seating[nx][ny]) == "#" {
				occupied_seats++
				break
			}
			nx = nx + derivative[0]
			ny = ny + derivative[1]
		}
	}
	return occupied_seats
}
func logic(seating []string, get_occupied_seats fn, number_of_empty_seats int) int {
	new_seating := []string{}
	changed := false
	total_occupied_seats := 0
	for i, row := range seating {
		new_row := ""
		for j, seat := range row {
			occupied_seats := get_occupied_seats(i, j, seating)
			// rules
			if string(seat) == "#" && occupied_seats >= number_of_empty_seats {
				new_row += "L"
				changed = true
			} else if string(seat) == "L" && occupied_seats == 0 {
				new_row += "#"
				changed = true
			} else {
				new_row += string(seat)
			}
			if string(new_row[j]) == "#" {
				total_occupied_seats++
			}
		}
		new_seating = append(new_seating, new_row)
	}
	if !changed {
		return total_occupied_seats
	} else {
		return logic(new_seating, get_occupied_seats, number_of_empty_seats)
	}

}
func main() {
	//ans := logic(getInput("test_input.txt"), get_adjacent_occupied_seats, 4)
	// fmt.Printf("Part 1: %d\n", ans)
	// ans = logic(getInput("test_input.txt"), get_visibly_occupied_seats, 5)
	// fmt.Printf("Part 2: %d\n", ans)

	input := getInput("input.txt")
	ans := logic(input, get_adjacent_occupied_seats, 4) // part1
	fmt.Printf("Part 1: %d\n", ans)
	ans = logic(input, get_visibly_occupied_seats, 5) // part2
	fmt.Printf("Part 2: %d\n", ans)

}
