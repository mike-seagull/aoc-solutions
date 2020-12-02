#!/usr/bin/env python


def part1():
    input_file = open("part1_input")
    puzzle_input = input_file.readline()

    i = 0
    consecutive_chars = 0
    total_sum = 0

    while i < len(puzzle_input):
        char = puzzle_input[i]
        if i == 0:
            if char == puzzle_input[len(puzzle_input) - 1]:  # compare first to last
                consecutive_chars = consecutive_chars + 1
        else:
            if char == puzzle_input[i - 1]:
                consecutive_chars = consecutive_chars + 1
            else:
                total_sum = total_sum + (consecutive_chars * int(puzzle_input[i - 1]))
                consecutive_chars = 0

        if i == len(puzzle_input) - 1 and consecutive_chars > 0:  # add the last if we need to
            total_sum = total_sum + (consecutive_chars * int(char))
        i = i + 1

    return total_sum


def part2():
    input_file = open("part2_input")
    puzzle_input = input_file.readline()
    first_half, second_half = puzzle_input[:len(puzzle_input)/2], puzzle_input[len(puzzle_input)/2:]  # split the string in half
    total_sum = 0
    for char1, char2 in zip(first_half, second_half):  # look at each half together
        if char1 == char2:
            total_sum = total_sum + (int(char1) * 2)
    return total_sum


print "part 1 = %i" % part1()
print "part 2 = %i" % part2()
